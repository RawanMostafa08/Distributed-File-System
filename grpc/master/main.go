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
)

type DataNode struct {
	IP              string
	Port            string
	NodeID          string
	IsDataNodeAlive bool
	HeartBeat       int
}

type FileData struct {
	FileID   string
	Filename string
	FilePath string
	FileSize int64
	NodeID   string
}

type LookUpTableTuple struct {
	File []FileData
}

type textServer struct {
	pb.UnimplementedDFSServer
}
type HeartBeatServer struct {
	pbHeartBeats.UnimplementedHeartbeatServiceServer
}

var lookupTuple LookUpTableTuple
var dataNodes []DataNode

var lookupTable = []FileData{
	{
		FileID:   "file_1",
		Filename: "data1.txt",
		FilePath: "/data/files/data1.txt",
		FileSize: 1024,
		NodeID:   "Node_1",
	},
	{
		FileID:   "file_2",
		Filename: "data2.txt",
		FilePath: "/data/files/data2.txt",
		FileSize: 2048,
		NodeID:   "Node_2",
	},
}

func getNodeByID(nodeID string) (DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return DataNode{}, errors.New("Node not found")
}

func (s *HeartBeatServer) KeepAlive(ctx context.Context, req *pbHeartBeats.HeartbeatRequest) (*pbHeartBeats.Empty, error) {
    fmt.Printf("Received heartbeat from: %s\n", req.NodeId)
    for i := range dataNodes {
        if dataNodes[i].NodeID == req.NodeId {
            dataNodes[i].HeartBeat += 1
            fmt.Printf("Node %s heartbeat count: %d\n", dataNodes[i].NodeID, dataNodes[i].HeartBeat)
        }
    }
    return &pbHeartBeats.Empty{}, nil
}

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes []string
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

	// for _, file := range lookupTuple.File {
	// 	fmt.Println(file.Filename + " " + req.GetFileName())
	// 	if file.Filename == req.GetFileName() {
	// 		if file.Node.IsDataNodeAlive == true {
	// 			nodes = append(nodes, file.Node.DataKeeperNode)
	// 			file_size = file.FileSize
	// 		}
	// 	}
	// }

	return &pb.DownloadPortsResponseBody{Addresses: nodes, FileSize: file_size}, nil
}

// Monitor node statuses and update lookup table
func monitorNodes() {
    for {
        time.Sleep(10 * time.Second)
        for i := range dataNodes {
            if dataNodes[i].HeartBeat == 0 {
                dataNodes[i].IsDataNodeAlive = false
                fmt.Printf("Node %s is dead\n", dataNodes[i].NodeID)
            } else {
                dataNodes[i].IsDataNodeAlive = true
                fmt.Printf("Node %s is alive (Heartbeats: %d)\n", dataNodes[i].NodeID, dataNodes[i].HeartBeat)
            }
            dataNodes[i].HeartBeat = 0 // Reset after checking
        }
    }
}

func main() {

	fmt.Println("master started...")
	var masterAddress, clientAddress string
	nodes := []string{}
	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)
	fmt.Printf("MASTER ADDRESS: '%s'\n", masterAddress)  // Should be ":8080"
    fmt.Printf("NODES: %v\n", nodes)                  
	for i, node := range nodes {
		dataNodes = append(dataNodes, DataNode{Port: node, NodeID: fmt.Sprintf("Node_%d", i), IsDataNodeAlive: true, HeartBeat: 1})
		// print("///////////",dataNodes[0].NodeID)
	}
	// check if nodes are alive and update the dataNodes list
	masterAddress = ":8081"
	lis, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	pbHeartBeats.RegisterHeartbeatServiceServer(s, &HeartBeatServer{})
	go monitorNodes()
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
