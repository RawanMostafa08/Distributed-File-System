package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"path/filepath"
	"sync"
	"time"

	// "strings"
	"strings"

	// "sync"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"

	pbHeartBeats "github.com/RawanMostafa08/Distributed-File-System/grpc/HeartBeats"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"

	// pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
	"github.com/RawanMostafa08/Distributed-File-System/grpc/models"
	pb_r_utils "github.com/RawanMostafa08/Distributed-File-System/grpc/replicate_utils"
)

type textServer struct {
	pb.UnimplementedDFSServer
	clientAddress string
}
type HeartBeatServer struct {
	pbHeartBeats.UnimplementedHeartbeatServiceServer
}

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

var dataNodes []models.DataNode
var lookupTable []models.FileData
var lookupTableMutex sync.Mutex

func getNodeByID(nodeID string) (models.DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return models.DataNode{}, errors.New("node not found")
}

func (s *textServer) UploadPortsRequest(ctx context.Context, req *pb.UploadRequestBody) (*pb.UploadResponseBody, error) {
	fmt.Println("1.Master received upload request")
	selectedNode := models.DataNode{IsDataNodeAlive: false}
	selectedPort := ""
	for _, node := range dataNodes {
		if node.IsDataNodeAlive && len(node.Port) > 0 {
			selectedNode = node
			randomIndex := rand.Intn(len(node.Port))
			selectedPort = node.Port[randomIndex]
			break
		}
	}
	if !selectedNode.IsDataNodeAlive || selectedPort == "" {
		return &pb.UploadResponseBody{
			DataNode_IP:  "",
			SelectedPort: selectedPort,
		}, fmt.Errorf("no alive data nodes or no free port found")
	}

	return &pb.UploadResponseBody{
		DataNode_IP:  selectedNode.IP,
		SelectedPort: selectedPort,
	}, nil
}

func (s *textServer) NodeMasterAckRequestUpload(ctx context.Context, req *pb.NodeMasterAckRequestBodyUpload) (*pb.Empty, error) {
	newFile := models.FileData{
		Filename: req.FileName,
		FilePath: req.FilePath,
		NodeID:   req.NodeId,
		FileSize: req.FileSize,
	}
	
	lookupTableMutex.Lock()
	lookupTable = append(lookupTable, newFile)
	lookupTableMutex.Unlock()
	fmt.Printf("4,5. Master notified and added file to lookup table: %s on node %s\n", req.FileName, req.DataNodeAddress)

	conn, err := grpc.Dial(s.clientAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		fmt.Println("Failed to connect to client:", err)
		return &pb.Empty{}, nil
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)

	_, err = c.MasterClientAckRequestUpload(context.Background(), &pb.MasterClientAckRequestBodyUpload{
		Message: fmt.Sprintf("Upload of file %s was successful.", req.FileName),
	})
	if err != nil {
		fmt.Println("Failed to send ack to client:", err)
	}

	return &pb.Empty{}, nil
}

func (s *textServer) MasterClientAckRequestUpload(ctx context.Context, req *pb.MasterClientAckRequestBodyUpload) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes []string
	var paths []string
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName())
	fmt.Println(req.GetFileName())
	lookupTableMutex.Lock()
	for _, file := range lookupTable {
		fmt.Println(file.Filename + " " + req.GetFileName())
		if file.Filename == req.GetFileName() {
			filenode, err := getNodeByID(file.NodeID)
			if err != nil {
				fmt.Println("Error getting node by ID:", err)

			} else if filenode.IsDataNodeAlive && len(filenode.Port) > 0 {
				paths = append(paths, file.FilePath)
				randomIndex := rand.Intn(len(filenode.Port))
				nodes = append(nodes, fmt.Sprintf("%s:%s", filenode.IP, filenode.Port[randomIndex]))
				file_size = file.FileSize
			}
		}
	}
	lookupTableMutex.Unlock()

	return &pb.DownloadPortsResponseBody{Addresses: nodes, Paths: paths, FileSize: file_size}, nil
}

func (s *HeartBeatServer) KeepAlive(ctx context.Context, req *pbHeartBeats.HeartbeatRequest) (*pbHeartBeats.Empty, error) {
	for i := range dataNodes {
		if dataNodes[i].NodeID == req.NodeId {
			dataNodes[i].HeartBeat += 1
		}
	}
	return &pbHeartBeats.Empty{}, nil
}

func ReplicateFile() {
	for {
		fmt.Println(lookupTable)
		time.Sleep(10 * time.Second)
		fmt.Println("Replicating Files")
		for _, file := range lookupTable {
			// get all nodes that have this file
			nodes := pb_r_utils.GetFileNodes(file.Filename, lookupTable, dataNodes)
			srcFile, err := pb_r_utils.GetSrcFileInfo(file, nodes)
			if err != nil {
				fmt.Println("Error getting source node ID:", err)
			} else {
				for len(nodes) < 3 && len(nodes) > 0 {
					valid, validPort, err := pb_r_utils.SelectNodeToCopyTo(nodes, dataNodes)
					if err != nil {
						fmt.Println(err)
						break
					} else {
						err = pb_r_utils.CopyFileToNode(srcFile, valid, validPort, &dataNodes)
						if err != nil {
							fmt.Println("Error copying file", err)
						} else {
							// update the lookup table
							lookupTableMutex.Lock()
							path := filepath.Join("files", valid)
							file := models.FileData{
								Filename: srcFile.Filename,
								FilePath: path,
								FileSize: srcFile.FileSize,
								NodeID:   valid}
							lookupTable = append(lookupTable, file)
							lookupTableMutex.Unlock()

						}
					}
					lookupTableMutex.Lock()
					nodes = pb_r_utils.GetFileNodes(file.Filename, lookupTable, dataNodes)
					lookupTableMutex.Unlock()
				}
			}
		}

	}
}

func cleaningLookuptable(nodeID string) {
	lookupTableMutex.Lock()
	defer lookupTableMutex.Unlock()

	newLookupTable := []models.FileData{}

	for _, file := range lookupTable {
		if file.NodeID != nodeID {
			newLookupTable = append(newLookupTable, file)
		} else {
			fmt.Printf("Removed file %s from lookup table (Node %s is dead)\n", file.Filename, nodeID)
		}
	}

	lookupTable = newLookupTable
}

// Monitor node statuses and update lookup table
func monitorNodes() {
	for {
		time.Sleep(5 * time.Second)
		for i := range dataNodes {
			if dataNodes[i].HeartBeat == 0 {
				dataNodes[i].IsDataNodeAlive = false
				fmt.Printf(" %s is dead\n", dataNodes[i].NodeID)
				cleaningLookuptable(dataNodes[i].NodeID)
			} else {
				dataNodes[i].IsDataNodeAlive = true
				fmt.Printf("Node %s is alive (Heartbeats: %d)\n", dataNodes[i].NodeID, dataNodes[i].HeartBeat)
			}
			dataNodes[i].HeartBeat = 0
		}
	}

}

func main() {

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	for i, node := range nodes {
		parts := strings.Split(node, ":")
		ip := parts[0]
		ports := strings.Split(parts[1], ",")
		dataNodes = append(dataNodes, models.DataNode{IP: ip, Port: ports, NodeID: fmt.Sprintf("Node_%d", i), IsDataNodeAlive: false, HeartBeat: 0})
	}

	lis, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	go ReplicateFile()

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxSendMsgSize(1024*1024*200), // 200MB send

	)
	pb.RegisterDFSServer(s, &textServer{
		clientAddress: clientAddress,
	})
	pbHeartBeats.RegisterHeartbeatServiceServer(s, &HeartBeatServer{})
	go monitorNodes()
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
