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

	// pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
	pb_r_utils "github.com/RawanMostafa08/Distributed-File-System/grpc/replicate_utils"
	"github.com/RawanMostafa08/Distributed-File-System/grpc/models"
)

// type DataNode struct {
// 	IP              string
// 	Port            string
// 	NodeID          string
// 	IsDataNodeAlive bool
// }

// type FileData struct {
// 	FileID   string
// 	Filename string
// 	FilePath string
// 	FileSize int64
// 	NodeID   string
// }

type LookUpTableTuple struct {
	File []models.FileData
}

type textServer struct {
	pb.UnimplementedDFSServer
}

// var lookupTable [] FileData
var dataNodes []models.DataNode

var lookupTable = []models.FileData{
	{
		FileID:   "file_1",
		Filename: "file1.mp4",
		FilePath: "files/Node_1",
		FileSize: 1055736,
		NodeID:   "Node_1",
	},
	{
		FileID:   "file_2",
		Filename: "file2.mp4",
		FilePath: "files/Node_2",
		FileSize: 1055736,
		NodeID:   "Node_2",
	},
}

func getNodeByID(nodeID string) (models.DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return models.DataNode{}, errors.New("Node not found")
}

func (s *textServer) DownloadPortsRequest(ctx context.Context, req *pb.DownloadPortsRequestBody) (*pb.DownloadPortsResponseBody, error) {
	fmt.Println("DownloadPortsRequest called")
	var nodes []string
	var paths []string
	var file_size int64
	file_size = 0
	fmt.Println(req.GetFileName())
	// ##########################################################################
	// Dummy Table Should be removed later , added id for each file
	// lookupTuple := LookUpTableTuple{
	// 	File: []models.FileData{
	// 		{FileID: "1", Filename: "file1.mp4", FilePath: "files/file1.mp4", FileSize: 1055736, NodeID: "Node_1"},
	// 		{FileID: "2", Filename: "file1.mp4", FilePath: "files/file1.mp4", FileSize: 1055736, NodeID: "Node_2"},
	// 	},
	// }
	// ##########################################################################

	for _, file := range lookupTable {
		fmt.Println(file.Filename + " " + req.GetFileName())
		if file.Filename == req.GetFileName() {
			filenode, err := getNodeByID(file.NodeID)
			if err != nil {
				fmt.Println("Error getting node by ID:", err)
			} else if filenode.IsDataNodeAlive == true {
				paths = append(paths, file.FilePath)
				nodes = append(nodes, filenode.Port)
				file_size = file.FileSize
			}
		}
	}

	return &pb.DownloadPortsResponseBody{Addresses: nodes, Paths: paths, FileSize: file_size}, nil
}

// // Replicate Helper Functions
// func getFileNodes(fileID string) []string {
// 	nodes := []string{}
// 	for _, f := range lookupTable {
// 		filenode, err := getNodeByID(f.NodeID)
// 		if err != nil {
// 			fmt.Println("Error getting node by ID:", err)
// 		} else if f.FileID == fileID && filenode.IsDataNodeAlive == true {
// 			nodes = append(nodes, filenode.NodeID)
// 		}
// 	}
// 	return nodes
// }

// func getSrcFileInfo(file FileData, nodes []string) (FileData, error) {
// 	for _, node := range nodes {
// 		if node == file.NodeID {
// 			file = FileData{
// 				FileID:   file.FileID,
// 				Filename: file.Filename,
// 				FilePath: file.FilePath,
// 				FileSize: file.FileSize,
// 				NodeID:   file.NodeID,
// 			}
// 			return file, nil
// 		}
// 	}
// 	return FileData{}, errors.New("Node not found")
// }

// func selectNodeToCopyTo(fileID string, fileNodes []string) (string, error) {
// 	// alive node , not in the list of nodes that have the file
// 	validNodes := []string{}
// 	for _, node := range dataNodes {
// 		flag := false
// 		if node.IsDataNodeAlive == true {
// 			for _, fileNode := range fileNodes {
// 				if fileNode == node.NodeID {
// 					flag = true
// 					break
// 				}
// 			}
// 			if flag == false {
// 				validNodes = append(validNodes, node.NodeID)
// 			}
// 		}
// 	}
// 	if len(validNodes) == 0 {
// 		return "", errors.New("No valid nodes to copy to")
// 	} else {
// 		return validNodes[0], nil
// 	}

// }

// func copyFileToNode(srcFile FileData, destNodeID string) error {
// 	srcNodeID := srcFile.NodeID
// 	srcNode, err := getNodeByID(srcNodeID)
// 	destNode, err := getNodeByID(destNodeID)
// 	if err != nil {
// 		return fmt.Errorf("Error getting node by ID: ", err)
// 	}
// 	conn, err := grpc.Dial(fmt.Sprintf("localhost%s", srcNode.Port), grpc.WithInsecure())
// 	if err != nil {
		
// 		return fmt.Errorf("Error in dial: ", err)
// 	}
	
// 	defer conn.Close()
// 	c := pb_r.NewDFSClient(conn)
// 	res, err := c.CopyNotification(context.Background(), &pb_r.CopyNotificationRequest{IsSrc: false, FileName: srcFile.Filename, FilePath: srcFile.FilePath, DestId: destNodeID, DestIp: destNode.IP, DestPort: destNode.Port})
// 	if err != nil {
// 		return fmt.Errorf("Error in CopyNotification: ", err)
// 	}
// 	fmt.Println("CopyNotification response:", res.Ack)
// 	if res.Ack != "Ack" {
// 		return fmt.Errorf(res.Ack)
// 	}
// 	return nil
// }

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
			nodes := pb_r_utils.GetFileNodes(file.FileID,lookupTable, dataNodes)
			srcFile, err := pb_r_utils.GetSrcFileInfo(file, nodes)
			if err != nil {
				fmt.Println("Error getting source node ID:", err)
			} else {
				for len(nodes) < 3 && len(nodes) > 0 {
					valid, err := pb_r_utils.SelectNodeToCopyTo(file.FileID,nodes, dataNodes)
					if err != nil {
						fmt.Println("Error selecting node to copy to:", err)
					} else {
						err = pb_r_utils.CopyFileToNode(srcFile, valid,dataNodes)
						if err != nil {
							fmt.Println("Error copying file", err)
						} else {
						// update the lookup table
						file := models.FileData{
							FileID:   srcFile.FileID,
							Filename: srcFile.Filename,
							FilePath: srcFile.FilePath,
							FileSize: srcFile.FileSize,
							NodeID:   valid}
						lookupTable = append(lookupTable, file)
						}
					}
					nodes = pb_r_utils.GetFileNodes(file.FileID,lookupTable, dataNodes)
				}
			}
		}

	}
}

func main() {

	fmt.Println("master started...")
	var masterAddress, clientAddress string
	nodes := []string{}
	pbUtils.ReadFile(&masterAddress, &clientAddress, &nodes)

	for i, node := range nodes {
		dataNodes = append(dataNodes, models.DataNode{IP:"localhost", Port: node, NodeID: fmt.Sprintf("Node_%d", i), IsDataNodeAlive: true})
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
