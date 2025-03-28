package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

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
	FileID int
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
var dataNodes [] DataNode 

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes [] string 
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName()) 
	// ##########################################################################
	// Dummy Table Should be removed later , added id for each file
	lookupTuple := LookUpTableTuple{
		File: []FileData{
			{FileID: 1, Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4" , FileSize:1055736 , Node: DataNode{DataKeeperNode: ":3000", IsDataNodeAlive: true}},
			{FileID: 2, Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize:1055736 , Node: DataNode{DataKeeperNode: ":8090", IsDataNodeAlive: true}},
		},
	}
	// ##########################################################################

	for _, file := range lookupTuple.File {
		fmt.Println(file.Filename + " " + req.GetFileName()) 
		if file.Filename == req.GetFileName() {
			if file.Node.IsDataNodeAlive == true {
				nodes = append(nodes, file.Node.DataKeeperNode)
				file_size = file.FileSize
			}
		}
	}

	return &pb.DownloadPortsResponseBody{Addresses: nodes, FileSize: file_size }, nil
}

// Replicate Helper Functions
func getFileNodes(fileID int)  []string {

	nodes := []string{}
	for _, f := range lookupTuple.File {
		if f.FileID == fileID && f.Node.IsDataNodeAlive == true {
			nodes = append(nodes, f.Node.DataKeeperNode )
		}
	}
	return  nodes
}

func selectNodeToCopyTo(fileID int, fileNodes []string) (string,error) {
	// alive node , not in the list of nodes that have the file
	validNodes := []string{}
	for _,node := range dataNodes {
		if node.IsDataNodeAlive == true  {
			for _, fileNode := range fileNodes {
				if fileNode == node.DataKeeperNode {
					continue
				}
				validNodes = append(validNodes, node.DataKeeperNode)
			}

		}
	}
	if len(validNodes) == 0 {
		return "",errors.New("No valid nodes to copy to")
	}
	return validNodes[0],nil
	

}

func CopyFileToNode(file FileData, srcNode string, destNode string) error {
	//notify machines
	//copy file
	//update lookup table 
	
	for _,f := range lookupTuple.File {
		if f.FileID == file.FileID {
			f.Node.DataKeeperNode = destNode
		}
	}
	return nil
}


// assume that each file record has fileid can be replicated in lookup table
// fileid 1, file1, node1
// fileid 1, file1, node2
// fileid 2, file1, node3
// fileid 2, file1, node1
func ReplicateFile() {
	for {
		time.Sleep(10 * time.Second)
		for _, file := range lookupTuple.File {
			// get all nodes that have this file
			nodes := getFileNodes(file.FileID);
			if len(nodes) < 3 && len(nodes) > 0{
					valid, err :=selectNodeToCopyTo(file.FileID,nodes)
					if err != nil {
						fmt.Errorf("Error selecting node to copy to")
					}
					fmt.Println("Selected Node to copy to: ", valid)
					err = CopyFileToNode(file, nodes[0], valid)
					if err != nil {
						fmt.Errorf("Error copying file")
					}
			}
		}

	}
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

	go ReplicateFile()

	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	fmt.Println("Server started. Listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
	
	

}
