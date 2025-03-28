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
	pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

type DataNode struct {
	IP string
	Port string
	NodeID string
	IsDataNodeAlive bool
}

type FileData struct {
	FileID string
	Filename string
	FilePath string
	FileSize int64
	NodeID string
}

type LookUpTableTuple struct {
	File []FileData
}




type textServer struct {
	pb.UnimplementedDFSServer
}
// var lookupTable [] FileData
var dataNodes [] DataNode 

var lookupTable = []FileData{
	{
		FileID:   "file_1",
		Filename: "data1.txt",
		FilePath: "/data/files/data1.txt",
		FileSize: 1024,
		NodeID:"Node_1",

	},
	{
		FileID:   "file_2",
		Filename: "data2.txt",
		FilePath: "/data/files/data2.txt",
		FileSize: 2048,
		NodeID: "Node_2",
	},
}

func getNodeByID(nodeID string) (DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return DataNode{}, errors.New("Node not found")
}


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
			{FileID: "1", Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4" , FileSize:1055736 , NodeID:"Node_1"},
			{FileID: "2", Filename: "file1.mp4", FilePath: "grpc\\files\\file1.mp4", FileSize:1055736 , NodeID: "Node_2"},
		},
	}
	// ##########################################################################

	for _, file := range lookupTuple.File {
		fmt.Println(file.Filename + " " + req.GetFileName()) 
		if file.Filename == req.GetFileName() {
			filenode,err:=getNodeByID(file.NodeID)
			if err != nil {
				fmt.Println("Error getting node by ID:", err)
			}else if filenode.IsDataNodeAlive == true {
				nodes = append(nodes, filenode.Port)
				file_size = file.FileSize
			}
		}
	}

	return &pb.DownloadPortsResponseBody{Addresses: nodes, FileSize: file_size }, nil
}

// Replicate Helper Functions
func getFileNodes(fileID string)  []string {
	nodes := []string{}
	for _, f := range lookupTable {
		filenode,err:=getNodeByID(f.NodeID)
		if err != nil {
			fmt.Println("Error getting node by ID:", err)
		}else if f.FileID == fileID && filenode.IsDataNodeAlive == true {
			nodes = append(nodes, filenode.NodeID)
		}
	}
	return  nodes
}

func selectNodeToCopyTo(fileID string, fileNodes []string) (string,error) {
	// alive node , not in the list of nodes that have the file
	validNodes := []string{}
	for _,node := range dataNodes {
		flag := false
		if node.IsDataNodeAlive == true  {
			for _, fileNode := range fileNodes {
				if fileNode == node.NodeID {
				flag = true
				break
				}
			}
			if flag == false {
			validNodes = append(validNodes, node.NodeID)
			}
		}
	}
	if len(validNodes) == 0 {
		return "",errors.New("No valid nodes to copy to")
	}else{
	return validNodes[0],nil
	}

}

func notify(file FileData, nodeID string,isSrc bool ) {	
	srcNode,err:=getNodeByID(nodeID)
	if err != nil {
		fmt.Println("Error getting node by ID:", err)
		return
	}
	conn, err := grpc.NewClient(fmt.Sprintf("localhost%s", srcNode.Port))
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()
	c := pb_r.NewDFSClient(conn)
	res,err:=c.CopyNotification(context.Background(), &pb_r.CopyNotificationRequest{is_src: isSrc, file_id: file.FileID})	
	if err != nil {
		fmt.Println("Error in CopyNotification:", err)
		return 
	}
	fmt.Println("CopyNotification response:", res.Ack)
}


func CopyFileToNode(file FileData, srcNodeID string, destNodeID string) error {
	//notify machines using different threads
	fmt.Println("Copying ",file.FileID ," from ",srcNodeID, " to ", destNodeID)
	go notify(file, srcNodeID, true)
	go notify(file, destNodeID, false)



	
	//copy file to destination node
	//update lookup table 
	file= FileData{
		FileID:   file.FileID,
		Filename: file.Filename,
		FilePath: file.FilePath,
		FileSize: file.FileSize,
		NodeID: destNodeID,}
	lookupTable = append(lookupTable, file)
	return nil
}


// assume that each file record has fileid can be replicated in lookup table
// fileid 1, file1, node1
// fileid 1, file1, node2
// fileid 2, file1, node3
// fileid 2, file1, node1
func ReplicateFile() {
	for {
		fmt.Println(lookupTable)
		time.Sleep(1 * time.Second)
		fmt.Println("Replicating Files")
		for _, file := range lookupTable {
			// get all nodes that have this file
			nodes := getFileNodes(file.FileID);
			for len(nodes) < 3 && len(nodes) > 0{
					valid, err :=selectNodeToCopyTo(file.FileID,nodes)
					if err != nil {
						fmt.Println("Error selecting node to copy to:", err)
					}else{
					err = CopyFileToNode(file, nodes[0], valid)
					if err != nil {
						fmt.Errorf("Error copying file")
					}
				}
				nodes = getFileNodes(file.FileID);
			}
		}

	}
}

func main() {

	
	fmt.Println("master started...")
	var masterAddress, clientAddress string
	nodes := []string{}
	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	
	for i, node := range nodes {
		dataNodes = append(dataNodes, DataNode{Port: node, NodeID: fmt.Sprintf("Node_%d",i), IsDataNodeAlive: true})
	}
	// check if nodes are alive and update the dataNodes list
	
	
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
