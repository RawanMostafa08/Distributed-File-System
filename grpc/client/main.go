package main

import (
	"context"
	"fmt"

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

func requestDownloadFile(nodeAddress string , file_name string) (*pb.DownloadFileResponseBody, error) {
	conn, err := grpc.Dial(nodeAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewDFSClient(conn)
	fmt.Println("Connected to Node",c)
		
	// Call the RPC method
	resp, err := c.DownloadFileRequest(context.Background(), &pb.DownloadFileRequestBody{FileName: file_name})

	if err != nil {
		fmt.Println("Error calling DownloadFileRequest:", err)
		return nil, err
	}
	return resp,nil	
	}

func main() {

	var masterAddress, clientAddress string
	nodes := []string{}

	// Read input from user
	fmt.Print("Enter File Name: ")
	var file_name string
	fmt.Scanln(&file_name)
		

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	// Download Logic
	// Connect to Master to get download ports
	var resp *pb.DownloadPortsResponseBody 
	var err error
	resp , err = requestDownloadPorts(masterAddress , file_name)		
	if err != nil {
		fmt.Println("Error calling DownloadPortsRequest:", err)
		return
	}
	fmt.Println("Nodes Master:", resp.Addresses)
	
	var resp_ *pb.DownloadFileResponseBody 
	for i := 0; i < len(resp.Addresses); i++ {
		resp_ , err = requestDownloadFile(resp.Addresses[i],file_name)
		if err != nil {
			fmt.Println("Error calling DownloadFileRequest:", err)
			return
		}
		fmt.Println("File Data:", resp_)
	}
}
