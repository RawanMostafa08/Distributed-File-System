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

// type DataNode struct {
// 	DataKeeperNode string
// 	IsDataNodeAlive bool
// }

// type FileData struct {
// 	Filename string
// 	FilePath string
// }

// type LookUpTableTuple struct {
// 	Node DataNode
// 	File []FileData
// }

type textServer struct {
	pb.UnimplementedDFSServer
}

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	return &pb.DownloadPortsResponseBody{Addresses: []string{":3000",":8090"}}, nil
}

func main() {
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
