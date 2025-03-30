package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	// "sync"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	"google.golang.org/grpc"
)

type DataNode struct {
	IP              string
	Port            string
	NodeID          string
	IsDataNodeAlive bool
}

type FileData struct {
	FileID   string
	Filename string
	FilePath string
	FileSize int64
	NodeID   string
}

var dataNodes []DataNode
var lookupTable []FileData

type textServer struct {
	pb.UnimplementedDFSServer
	clientAddress string
}

func getNodeByID(nodeID string) (DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return DataNode{}, errors.New("node not found")
}

func (s *textServer) UploadPortsRequest(ctx context.Context, req *pb.UploadRequestBody) (*pb.UploadResponseBody, error) {
	fmt.Println("1.Master received upload request")
	selectedNode := DataNode{IsDataNodeAlive: false}
	for _,node := range(dataNodes){
		if node.IsDataNodeAlive{
			selectedNode = node
			break
		}
	}

	if !selectedNode.IsDataNodeAlive{
		return &pb.UploadResponseBody{
			DataNode_IP:  selectedNode.IP,
			SelectedPort: selectedNode.Port,
		}, fmt.Errorf("no alive data nodes found")
	}

	return &pb.UploadResponseBody{
		DataNode_IP:  selectedNode.IP,
		SelectedPort: selectedNode.Port,
	}, nil
}

func (s *textServer) NodeMasterAckRequestUpload(ctx context.Context, req *pb.NodeMasterAckRequestBodyUpload) (*pb.Empty, error) {

	newFile := FileData{
		Filename: req.FileName,
		FilePath: req.FilePath,
		NodeID:   strconv.Itoa(int(req.NodeId)),
	}

	lookupTable = append(lookupTable, newFile)

	fmt.Printf("4,5. Master notified and added file to lookup table: %s on node %s\n", req.FileName, req.DataNodeAddress)
	conn, err := grpc.Dial(s.clientAddress, grpc.WithInsecure())
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
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName())
	// ##########################################################################
	// Dummy Table Should be removed later , added id for each file
	lookupTuple := []FileData{
		{FileID: "1", Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize: 1055736, NodeID: "Node_1"},
		{FileID: "2", Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize: 1055736, NodeID: "Node_2"},
	}
	// ##########################################################################

	for _, file := range lookupTuple {
		fmt.Println(file.Filename + " " + req.GetFileName())
		if file.Filename == req.GetFileName() {
			filenode, err := getNodeByID(file.NodeID)
			if err != nil {
				fmt.Println("Error getting node by ID:", err)
			} else if filenode.IsDataNodeAlive {
				nodes = append(nodes, filenode.Port)
				file_size = file.FileSize
			}
		}
	}

	return &pb.DownloadPortsResponseBody{Addresses: nodes, FileSize: file_size}, nil
}

func main() {

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	for i, node := range nodes {
		parts := strings.Split(node, ":")
		dataNodes = append(dataNodes, DataNode{IP: parts[0], Port: parts[1], NodeID: fmt.Sprintf("Node_%d", i), IsDataNodeAlive: true})
	}

	lis, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{
		clientAddress: clientAddress,
	})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
