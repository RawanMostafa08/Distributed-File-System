package main

import (
	"context"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload" // Import the generated package
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
)

func requestUploadPort(masterAddress string) (*pb.UploadResponseBody, error) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
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
	return resp, nil
}

func uploadFile(nodeAddress, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %v",err)
	}
	defer file.Close()
    fileName := filepath.Base(filePath)

	// Read file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file %v",err)
	}

	// Connect to data node
	conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)

	// Upload file to data node
	_, err = c.UploadFileRequest(context.Background(), &pb.UploadFileRequestBody{
		NodeAddress: nodeAddress,
		FileData:    fileData,
		FileName: fileName,
	})
	return err
}

func requestDownloadPorts(masterAddress string, file_name string) (*pb.DownloadPortsResponseBody, error) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
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

func requestDownloadFile(nodeAddress, fileName string, start, end int64, wg *sync.WaitGroup, chunks map[int][]byte, index int, mu *sync.Mutex) {
	defer wg.Done()
	conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect to node:", err)
		return
	}
	defer conn.Close()

	c := pb.NewDFSClient(conn)
	resp, err := c.DownloadFileRequest(context.Background(), &pb.DownloadFileRequestBody{FileName: fileName, Start: start, End: end})

	if err != nil {
		fmt.Println("Error downloading chunk:", err)
		return
	}

	mu.Lock()
	chunks[index] = resp.FileData
	mu.Unlock()

}

func downloadFile(Addresses []string, file_name string, fileSize int64) {
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
		go requestDownloadFile(node, file_name, start, end, &wg, chunks, i, &mu)
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

	fmt.Println("File downloaded successfully!")

}

func main() {

	// read parties addresses
	var masterAddress, clientAddress string
	nodes := []pbUtils.Node{}
	fmt.Println("Choose operation:")
	fmt.Println("1. Upload file")
	fmt.Println("2. Download file")
	var choice int
	fmt.Scanln(&choice)

	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

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
		fmt.Println(resp.DataNode_IP, resp.SelectedPort)

		// Upload file to selected data node
		dataNodeAddress := fmt.Sprintf("%s:%s", resp.DataNode_IP, resp.SelectedPort)
		err = uploadFile(dataNodeAddress, filePath)
		if err != nil {
			fmt.Println("Error uploading file:", err)
			return
		}

		fmt.Println("File uploaded successfully!")

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
		fmt.Println("Nodes Master:", resp.Addresses)

		fileSize := resp.FileSize

		downloadFile(resp.Addresses, file_name, fileSize)
	default:
		fmt.Println("Invalid choice")
	}

}
