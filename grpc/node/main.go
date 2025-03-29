package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

type textServer struct {
	pb.UnimplementedDFSServer
	masterAddress string
	nodeAddress string
	nodeID int32
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


func (s *textServer) UploadFileRequest(ctx context.Context, req *pb.UploadFileRequestBody) (*pb.Empty, error) {

	if err := os.MkdirAll("grpc/files", os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Save the file
	filePath := filepath.Join("grpc", "files", req.FileName)
	if err := ioutil.WriteFile(filePath, req.FileData, 0644); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}


	// Notify master tracker
	conn, err := grpc.Dial(s.masterAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to master: %v", err)
	}
	defer conn.Close()

	masterClient := pb.NewDFSClient(conn)
	_, err = masterClient.NodeMasterAckRequestUpload(ctx, &pb.NodeMasterAckRequestBodyUpload{
		FileName: req.FileName,
		FilePath: filePath,
		DataNodeAddress: s.nodeAddress,
		NodeId: s.nodeID,
		Status: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to notify master: %v", err)
	}

	fmt.Printf("File %s saved successfully and master notified\n", req.FileName)
	return &pb.Empty{}, nil

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

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	lis, err := net.Listen("tcp",nodes[node_index])
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{
		masterAddress: masterAddress,
		nodeAddress: nodes[node_index],
		nodeID: node_index,
		})
	fmt.Println("Server started. Listening on ", nodes[node_index], "...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
}
