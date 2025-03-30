package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"

	pbHeartBeats "github.com/RawanMostafa08/Distributed-File-System/grpc/HeartBeats"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"

	// pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
	pb_r_utils "github.com/RawanMostafa08/Distributed-File-System/grpc/replicate_utils"
	"github.com/RawanMostafa08/Distributed-File-System/grpc/models"
)


type LookUpTableTuple struct {
	File []models.FileData
}

type textServer struct {
	pb.UnimplementedDFSServer
}
type HeartBeatServer struct {
	pbHeartBeats.UnimplementedHeartbeatServiceServer
}

var lookupTuple LookUpTableTuple
var dataNodes []models.DataNode

var lookupTable = []models.FileData{
	{
		FileID:   "file_1",
		Filename: "file1.mp4",
		FilePath: "files/Node_1",
		FileSize: 1055736,
		NodeID:   "Node_1",
	},
	{
		FileID:   "file_2",
		Filename: "file2.mp4",
		FilePath: "files/Node_2",
		FileSize: 1055736,
		NodeID:   "Node_2",
	},
}

func getNodeByID(nodeID string) (models.DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return models.DataNode{}, errors.New("Node not found")
}

func (s *HeartBeatServer) KeepAlive(ctx context.Context, req *pbHeartBeats.HeartbeatRequest) (*pbHeartBeats.Empty, error) {
    for i := range dataNodes {
        if dataNodes[i].NodeID == req.NodeId {
            dataNodes[i].HeartBeat += 1
        }
    }
    return &pbHeartBeats.Empty{}, nil
}

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes []string
	var paths []string
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName())
	// ##########################################################################
	// Dummy Table Should be removed later
	// lookupTuple := LookUpTableTuple{
	// 	File: []FileData{
	// 		{Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4" , FileSize:1055736 , Node: DataNode{DataKeeperNode: ":3000", IsDataNodeAlive: true}},
	// 		{Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize:1055736 , Node: DataNode{DataKeeperNode: ":8090", IsDataNodeAlive: true}},
	// 	},
	// }
	// // ##########################################################################

	for _, file := range lookupTable {
		fmt.Println(file.Filename + " " + req.GetFileName())
		if file.Filename == req.GetFileName() {
			filenode, err := getNodeByID(file.NodeID)
			if err != nil {
				fmt.Println("Error getting node by ID:", err)
			} else if filenode.IsDataNodeAlive == true {
				paths = append(paths, file.FilePath)
				nodes = append(nodes, filenode.Port)
				file_size = file.FileSize
			}
		}
	}

	return &pb.DownloadPortsResponseBody{Addresses: nodes, Paths: paths, FileSize: file_size}, nil
}

// Monitor node statuses and update lookup table
func monitorNodes() {
    for {
        time.Sleep(5 * time.Second)
        for i := range dataNodes {
            if dataNodes[i].HeartBeat == 0 {
                dataNodes[i].IsDataNodeAlive = false
                fmt.Printf(" %s is dead\n", dataNodes[i].NodeID)
            } else {
                dataNodes[i].IsDataNodeAlive = true
                fmt.Printf("Node %s is alive (Heartbeats: %d)\n", dataNodes[i].NodeID, dataNodes[i].HeartBeat)
            }
            dataNodes[i].HeartBeat = 0 
        }
    }
	
}


// assume that each file record has fileid can be replicated in lookup table
// fileid 1, file1, node1
// fileid 1, file1, node2
// fileid 2, file1, node3
// fileid 2, file1, node1
func ReplicateFile() {
	for {
		fmt.Println(lookupTable)
		time.Sleep(10 * time.Second)
		fmt.Println("Replicating Files")
		for _, file := range lookupTable {
			// get all nodes that have this file
			nodes := pb_r_utils.GetFileNodes(file.FileID,lookupTable, dataNodes)
			srcFile, err := pb_r_utils.GetSrcFileInfo(file, nodes)
			if err != nil {
				fmt.Println("Error getting source node ID:", err)
			} else {
				for len(nodes) < 3 && len(nodes) > 0 {
					valid, err := pb_r_utils.SelectNodeToCopyTo(file.FileID,nodes, dataNodes)
					if err != nil {
						fmt.Println("Error selecting node to copy to:", err)
						break
					} else {
						err = pb_r_utils.CopyFileToNode(srcFile, valid,dataNodes)
						if err != nil {
							fmt.Println("Error copying file", err)
						} else {
						// update the lookup table
						file := models.FileData{
							FileID:   srcFile.FileID,
							Filename: srcFile.Filename,
							FilePath: srcFile.FilePath,
							FileSize: srcFile.FileSize,
							NodeID:   valid}
						lookupTable = append(lookupTable, file)
						}
					}
					nodes = pb_r_utils.GetFileNodes(file.FileID,lookupTable, dataNodes)
				}
			}
		}

	}
}

func main() {

	fmt.Println("master started...")
	var masterAddress, clientAddress string
	nodes := []string{}
	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	for i, node := range nodes {
		dataNodes = append(dataNodes, models.DataNode{IP:"localhost", Port: node, NodeID: fmt.Sprintf("Node_%d", i), IsDataNodeAlive: false , HeartBeat: 0})
	}
	// check if nodes are alive and update the dataNodes list

	lis, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	go ReplicateFile()

	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	pbHeartBeats.RegisterHeartbeatServiceServer(s, &HeartBeatServer{})
	go monitorNodes()
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
