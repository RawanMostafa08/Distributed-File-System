package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"

	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"

	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	"google.golang.org/grpc"

	pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
)

type textServer struct {
	pb.UnimplementedDFSServer
}

type replicateServer struct {
	pb_r.UnimplementedDFSServer
}

func (s *textServer) DownloadFileRequest(ctx context.Context, req *pb.DownloadFileRequestBody) (*pb.DownloadFileResponseBody, error) {
	fmt.Println("DownloadFileRequest called for:", req.FileName, "Range:", req.Start, "-", req.End)

	// Open the file
	filePath := "grpc\\files\\" + req.FileName
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get the requested chunk
	chunkSize := req.End - req.Start + 1
	data := make([]byte, chunkSize)

	// Seek to the start position
	_, err = file.Seek(req.Start, 0)
	if err != nil {
		return nil, err
	}

	// Read the requested bytes
	n, err := file.Read(data)
	if err != nil && err != io.EOF {
		return nil, err
	}

	// Trim the slice in case we read fewer bytes than requested
	data = data[:n]

	// Return the chunk as a response
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

func (s *replicateServer) CopyFile(ctx context.Context, req *pb_r.CopyFileRequest) (*pb_r.CopyFileResponse, error) {
	filePath := fmt.Sprintf("files/%s/%s", req.DestId, req.FileName)
	err := os.WriteFile(filePath, req.FileData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return nil, err
	}
	fmt.Println("File copied successfully to:", filePath)
	return &pb_r.CopyFileResponse{Ack: "ACK"}, nil
}

func (s *replicateServer) CopyNotification(ctx context.Context, req *pb_r.CopyNotificationRequest) (*pb_r.CopyNotificationResponse, error) {
	filePath := fmt.Sprintf("%s/%s", req.FilePath, req.FileName)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s%s", req.DestIp, req.DestPort), grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()

	c := pb_r.NewDFSClient(conn)
	res, err := c.CopyFile(context.Background(), &pb_r.CopyFileRequest{FileName: req.FileName, FileData: fileContent, DestId: req.DestId })
	if err != nil {
		fmt.Println("Error in CopyFile:", err)
		return nil, err
	}
	if res.Ack == "ACK" {
		return &pb_r.CopyNotificationResponse{Ack: "Ack"}, nil
	}
	return nil, errors.New("File not copied to destination node")

}

func main() {

	// Read input from user
	fmt.Print("Enter node index : ")
	var node_index int32
	fmt.Scanln(&node_index)

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	lis, err := net.Listen("tcp", nodes[node_index])
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	pb_r.RegisterDFSServer(s, &replicateServer{})
	fmt.Println("Server started. Listening on port ", nodes[node_index], "...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
