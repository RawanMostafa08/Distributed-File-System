package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload" // Import the generated package
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
)

type textClientServer struct{
	pb.UnimplementedDFSServer
}

func (s *textClientServer) MasterClientAckRequestUpload(ctx context.Context, req *pb.MasterClientAckRequestBodyUpload) (*pb.Empty, error) {
	fmt.Printf("6. Client notified-> %s",req.Message)
	return &pb.Empty{}, nil
}
func startClientServer(clientAddress string) {
	lis, err := net.Listen("tcp", clientAddress)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxSendMsgSize(1024*1024*200), // 200MB send	
	)
	pb.RegisterDFSServer(s, &textClientServer{})
	fmt.Printf("Client gRPC server listening at %s\n", clientAddress)
	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}

func requestUploadPort(masterAddress string) (*pb.UploadResponseBody, error) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)

	// Call the RPC method
	resp, err := c.UploadPortsRequest(context.Background(), &pb.UploadRequestBody{MasterAddress: masterAddress})
	if err != nil {
		fmt.Println("Error calling UploadRequest:", err)
		return nil, err
	}
	fmt.Printf("2. Client received node address %s:%s",resp.DataNode_IP, resp.SelectedPort)

	return resp, nil
}

func uploadFile(nodeAddress, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %v", err)
	}
	defer file.Close()
	fileName := filepath.Base(filePath)

	// Read file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file %v", err)
	}

	// Connect to data node
	// conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure())
	conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)

	// Upload file to data node
	_, err = c.UploadFileRequest(context.Background(), &pb.UploadFileRequestBody{
		NodeAddress: nodeAddress,
		FileData:    fileData,
		FileName:    fileName,
		FileSize: int64(len(fileData)),
	})
	return err
}

func requestDownloadPorts(masterAddress string, file_name string) (*pb.DownloadPortsResponseBody, error) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)
	fmt.Println("Connected to Master", c)

	// Call the RPC method
	resp, err := c.DownloadPortsRequest(context.Background(), &pb.DownloadPortsRequestBody{FileName: file_name})

	if err != nil {
		fmt.Println("Error calling DownloadPortsRequest:", err)
		return nil, err
	}
	return resp, nil

}

func requestDownloadFile(nodeAddress, fileName string,filePath string, start, end int64, wg *sync.WaitGroup, chunks map[int][]byte, index int, mu *sync.Mutex) {
	defer wg.Done()
	// conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure())
	conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		fmt.Println("Failed to connect to node:", err)
		return
	}
	defer conn.Close()

	c := pb.NewDFSClient(conn)
	resp, err := c.DownloadFileRequest(context.Background(), &pb.DownloadFileRequestBody{FileName: fileName, FilePath: filePath,Start: start, End: end})

	if err != nil {
		fmt.Println("Error downloading chunk:", err)
		return
	}

	mu.Lock()
	chunks[index] = resp.FileData
	mu.Unlock()

}

func downloadFile(Addresses []string,Paths []string ,file_name string, fileSize int64) {
	chunkSize := int64(math.Ceil(float64(fileSize) / float64(len(Addresses))))
	var wg sync.WaitGroup
	chunks := make(map[int][]byte)
	var mu sync.Mutex

	for i, node := range Addresses {
		start := int64(i) * chunkSize
		end := start + chunkSize - 1
		if end >= fileSize {
			end = fileSize - 1
		}
		wg.Add(1)
		go requestDownloadFile(node, file_name,Paths[i], start, end, &wg, chunks, i, &mu)
	}

	wg.Wait()

	// Reconstruct the file
	outputFile, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	for i := 0; i < len(Addresses); i++ {
		if data, exists := chunks[i]; exists {
			outputFile.Write(data)
		}
	}

}

func main() {
	// read parties addresses
	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)
	go startClientServer(clientAddress)

	fmt.Println("Choose operation:")
	fmt.Println("1. Upload file")
	fmt.Println("2. Download file")
	var choice int
	fmt.Scanln(&choice)


	switch choice {
	case 1:
		// Upload logic
		fmt.Print("Enter file path to upload: ")
		var filePath string
		fmt.Scanln(&filePath)

		// Get available data node from master
		resp, err := requestUploadPort(masterAddress)
		if err != nil {
			fmt.Println("Error getting upload port:", err)
			return
		}

		// Upload file to selected data node
		dataNodeAddress := fmt.Sprintf("%s:%s", resp.DataNode_IP, resp.SelectedPort)
		err = uploadFile(dataNodeAddress, filePath)
		if err != nil {
			fmt.Println("Error uploading file:", err)
			return
		}

	case 2:
		// Existing download logic
		fmt.Print("Enter File Name To Download: ")
		var file_name string
		fmt.Scanln(&file_name)

		// Connect to Master to get download ports
		var resp *pb.DownloadPortsResponseBody
		var err error
		resp, err = requestDownloadPorts(masterAddress, file_name)
		if err != nil {
			fmt.Println("Error calling DownloadPortsRequest:", err)
			return
		}
		if len(resp.Addresses) == 0{
			fmt.Print("file not found on any node")
			return
		}
		fmt.Println("Nodes Master:", resp.Addresses)

		fileSize := resp.FileSize

		downloadFile(resp.Addresses, resp.Paths,file_name, fileSize)
	default:
		fmt.Println("Invalid choice")
	}

}
