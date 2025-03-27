package main

import (
	"context"
	"fmt"
	"net"
	"io/ioutil"
	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

type textServer struct {
	pb.UnimplementedDFSServer
}

func (s *textServer) DownloadFileRequest(ctx context.Context, req *pb.DownloadFileRequestBody) (*pb.DownloadFileResponseBody, error) {
	fmt.Println("DownloadFileRequest called")
	var data []byte
	data , err := ReadMP4File("grpc\\files\\"+req.FileName)
	if err != nil {
		return nil, err
	}
	return &pb.DownloadFileResponseBody{FileData: data}, nil
}

// ReadMP4File reads an MP4 file and returns its content as a byte slice.
func ReadMP4File(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}



func main() {

	// Read input from user
	fmt.Print("Enter node index : ")
	var node_index int32
	fmt.Scanln(&node_index)
		

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	
	lis, err := net.Listen("tcp", nodes[node_index])
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	fmt.Println("Server started. Listening on port ",nodes[node_index],"...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}



}
