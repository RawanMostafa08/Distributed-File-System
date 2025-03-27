package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"os"
	// "net"
	"google.golang.org/grpc"
	// "google.golang.org/protobuf/internal/encoding/text"
	// ""
	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload" // Import the generated package
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	// "google.golang.org/grpc"
)

func requestDownloadPorts(masterAddress string , file_name string) (*pb.DownloadPortsResponseBody, error) {
	conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)
	fmt.Println("Connected to Master",c)

	// Call the RPC method
	resp, err := c.DownloadPortsRequest(context.Background(), &pb.DownloadPortsRequestBody{FileName: file_name})

	if err != nil {
		fmt.Println("Error calling DownloadPortsRequest:", err)
		return nil, err
	}
	return resp,nil

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

	var masterAddress, clientAddress string
	nodes := []string{}

	// Download Logic
	// Read input from user
	fmt.Print("Enter File Name To Download : ")
	var file_name string
	fmt.Scanln(&file_name)
		

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	// Connect to Master to get download ports
	var resp *pb.DownloadPortsResponseBody 
	var err error
	resp , err = requestDownloadPorts(masterAddress , file_name)		
	if err != nil {
		fmt.Println("Error calling DownloadPortsRequest:", err)
		return
	}
	fmt.Println("Nodes Master:", resp.Addresses)
	
	fileSize := resp.FileSize

	downloadFile(resp.Addresses, file_name, fileSize)
}
