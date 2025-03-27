package main

import (
	"context"
	"fmt"
	"net"
	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

type DataNode struct {
	DataKeeperNode string
	IsDataNodeAlive bool
}

type FileData struct {
	Filename string
	FilePath string
	FileSize int64
	Node DataNode
}

type LookUpTableTuple struct {
	File []FileData
}


type textServer struct {
	pb.UnimplementedDFSServer
}
var lookupTuple LookUpTableTuple

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes [] string 
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName()) 
	
	lookupTuple := LookUpTableTuple{
		File: []FileData{
			{Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4" , FileSize:1055736 , Node: DataNode{DataKeeperNode: ":3000", IsDataNodeAlive: true}},
			{Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize:1055736 , Node: DataNode{DataKeeperNode: ":8090", IsDataNodeAlive: true}},
		},
	}

	for _, file := range lookupTuple.File {
		fmt.Println(file.Filename + " " + req.GetFileName()) 
		if file.Filename == req.GetFileName() {
			if file.Node.IsDataNodeAlive == true {
				nodes = append(nodes, file.Node.DataKeeperNode)
				file_size = file.FileSize
			}
		}
	}
	fmt.Println(nodes) 

	return &pb.DownloadPortsResponseBody{Addresses: nodes, FileSize: file_size }, nil
}

func main() {

	
	fmt.Println("LookUpTableTuple Example:")
	for _, file := range lookupTuple.File {
		fmt.Println("File:", file.Filename, "| Path:", file.FilePath)
	}


	
	fmt.Println("master started...")

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	
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
