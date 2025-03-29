package main

import (
	"context"
	"errors"
	"fmt"
	"net"

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

	return &pb.UploadResponseBody{
		SelectedPort: dataNodes[0].Port,
		DataNode_IP:  dataNodes[0].IP,
	}, nil
}



func (s *textServer) NodeMasterAckRequest(ctx context.Context, req *pb.NodeMasterAckRequestBody) (*pb.Empty, error) {
	fmt.Println("ackkkkkk")

	return &pb.Empty{}, nil
}

//	func (s *textServer) MasterClientAckRequest(ctx context.Context, req *pb.MasterClientAckRequestBody) (*pb.Empty, error) {
//		// Here you would handle the client acknowledgment
//		return &pb.Empty{}, nil
//	}
func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes []string
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName())
	// ##########################################################################
	// Dummy Table Should be removed later , added id for each file
	lookupTuple :=  []FileData{
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

	fmt.Println("master started...")
	var masterAddress, clientAddress string
	nodes := []pbUtils.Node{}

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	for i, node := range nodes {
		dataNodes = append(dataNodes, DataNode{IP:node.IP ,Port: node.Port, NodeID: fmt.Sprintf("Node_%d",i), IsDataNodeAlive: true})
	}

	lis, err := net.Listen("tcp", masterAddress)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}

}
