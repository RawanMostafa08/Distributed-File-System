package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"path/filepath"
	"time"

	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"

	pbHeartBeats "github.com/RawanMostafa08/Distributed-File-System/grpc/HeartBeats"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"

	"google.golang.org/grpc"

	pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
)

type textServer struct {
	pb.UnimplementedDFSServer
	masterAddress string
	nodeAddress   string
	nodeID        string
}

type HeartBeatServer struct {
	pbHeartBeats.UnimplementedHeartbeatServiceServer
}

type replicateServer struct {
	pb_r.UnimplementedDFSServer
}

func ackToMaster(ctx context.Context, masterAddress string, nodeID string, port string) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*1024),
		grpc.MaxCallSendMsgSize(1024*1024*1024), 
	))
	if err != nil {
		return
	}
	defer conn.Close()

}

func (s *textServer) DownloadFileRequest(ctx context.Context, req *pb.DownloadFileRequestBody) (*pb.DownloadFileResponseBody, error) {
	fmt.Println("DownloadFileRequest called for:", req.FileName, "Range:", req.Start, "-", req.End)

	// Open the file
	// filePath := fmt.Sprintf("%s/%s", req.FilePath, req.FileName)
	parts := strings.FieldsFunc(req.FilePath, func(r rune) bool {
		return r == '/' || r == '\\'
	})
	normalizedPath := filepath.Join(parts...)
	fmt.Println("Normalized path:", normalizedPath)

	filePath := filepath.Join(normalizedPath, req.FileName)
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

	if err := os.MkdirAll("files", os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// Save the file
	filePath := filepath.Join("files", s.nodeID)

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	filePathU := filepath.Join("files", s.nodeID, req.FileName)

	if err := os.WriteFile(filePathU, req.FileData, 0644); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Notify master tracker
	conn, err := grpc.Dial(s.masterAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*1024), 
		grpc.MaxCallSendMsgSize(1024*1024*1024), 
	))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to master: %v", err)
	}
	defer conn.Close()

	masterClient := pb.NewDFSClient(conn)
	_, err = masterClient.NodeMasterAckRequestUpload(ctx, &pb.NodeMasterAckRequestBodyUpload{
		FileName:        req.FileName,
		FilePath:        filePath,
		DataNodeAddress: s.nodeAddress,
		NodeId:          s.nodeID,
		FileSize:        req.FileSize,
		Status:          true,
		ClientAddress: req.ClientAddress,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to notify master: %v", err)
	}

	fmt.Printf("3,4. File %s saved successfully\n", req.FileName)
	return &pb.Empty{}, nil

}

// ReadMP4File reads an MP4 file and returns its content as a byte slice.
func ReadMP4File(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func pingMaster(nodeIndex int32, masterAddress string) {
	for {
		conn, err := grpc.Dial(masterAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*1024), 
			grpc.MaxCallSendMsgSize(1024*1024*1024), 
		))
		if err != nil {
			fmt.Printf("Failed to connect to master: %v \n", err)
		}

		master := pbHeartBeats.NewHeartbeatServiceClient(conn)
		for {
			_, err := master.KeepAlive(context.Background(), &pbHeartBeats.HeartbeatRequest{
				NodeId: fmt.Sprintf("Node_%d", nodeIndex),
			})
			if err != nil {
				fmt.Printf("Heartbeat failed: %v \n", err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *replicateServer) CopyFile(ctx context.Context, req *pb_r.CopyFileRequest) (*pb_r.CopyFileResponse, error) {
	if err := os.MkdirAll("files", os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}
	folderPath := filepath.Join("files", req.DestId)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	filePath := filepath.Join(folderPath, req.FileName)
	err := os.WriteFile(filePath, req.FileData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return nil, err
	}

	fmt.Println("File copied successfully to:", req.DestId)
	return &pb_r.CopyFileResponse{Ack: "ACK"}, nil
}

func (s *replicateServer) CopyNotification(ctx context.Context, req *pb_r.CopyNotificationRequest) (*pb_r.CopyNotificationResponse, error) {
	filePath := filepath.Join(req.FilePath, req.FileName)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// conn, err := grpc.Dial(fmt.Sprintf("%s:%s", req.DestIp, req.DestPort), grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", req.DestIp, req.DestPort), grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*1024), 
		grpc.MaxCallSendMsgSize(1024*1024*1024), 
	))
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()

	c := pb_r.NewDFSClient(conn)
	res, err := c.CopyFile(context.Background(), &pb_r.CopyFileRequest{FileName: req.FileName, FileData: fileContent, DestId: req.DestId})
	if err != nil {
		fmt.Println("Error in CopyFile:", err)
		return nil, err
	}
	if res.Ack == "ACK" {
		fmt.Printf("Copy Notification sent to: %s \n",req.DestId)
		return &pb_r.CopyNotificationResponse{Ack: "Ack"}, nil
	}
	return nil, errors.New("file not copied to destination node")

}

func main() {

	// Read input from user
	fmt.Print("Enter node index : ")
	var node_index int32
	fmt.Scanln(&node_index)

	var masterAddress string
	nodes := []string{}

	pbUtils.ReadFile_node(&masterAddress, &nodes)
	parts := strings.Split(nodes[node_index], ":")
	ip := parts[0]
	node_ports := strings.Split(parts[1], ",")

	for _, port := range node_ports {
		go func(p string) {
			address := fmt.Sprintf(":%s", strings.TrimSpace(p))
			lis, err := net.Listen("tcp", address)
			if err != nil {
				fmt.Println("failed to listen on", address, ":", err)
				return
			}

			s := grpc.NewServer(
				grpc.MaxRecvMsgSize(1024*1024*1024), 
				grpc.MaxSendMsgSize(1024*1024*1024),
			)

			pb.RegisterDFSServer(s, &textServer{
				masterAddress: masterAddress,
				nodeAddress:   fmt.Sprintf("%s:%s", ip, p),
				nodeID:        fmt.Sprintf("Node_%d", node_index),
			})
			pb_r.RegisterDFSServer(s, &replicateServer{})

			fmt.Println("Server started. Listening on", address, "...")
			if err := s.Serve(lis); err != nil {
				fmt.Println("failed to serve on", address, ":", err)
			}
		}(port)
	}

	go pingMaster(node_index, masterAddress)
	select {}

}