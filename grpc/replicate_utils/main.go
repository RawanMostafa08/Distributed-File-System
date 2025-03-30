
package Replicate_utils

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
	"github.com/RawanMostafa08/Distributed-File-System/grpc/models"

)


// Replicate Helper Functions

func GetNodeByID(nodeID string, dataNodes []models.DataNode) (models.DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return models.DataNode{}, errors.New("Node not found")
}

func GetFileNodes(fileName string, lookupTable []models.FileData, dataNodes []models.DataNode) []string {
	nodes := []string{}
	for _, f := range lookupTable {
		filenode, err := GetNodeByID(f.NodeID,dataNodes)
		if err != nil {
			fmt.Println("Error getting node by ID:", err)
		} else if f.Filename == fileName && filenode.IsDataNodeAlive == true {
			nodes = append(nodes, filenode.NodeID)
		}
	}
	return nodes
}

func GetSrcFileInfo(file models.FileData, nodes []string) (models.FileData, error) {
	for _, node := range nodes {
		if node == file.NodeID {
			file = models.FileData{
				Filename: file.Filename,
				FilePath: file.FilePath,
				FileSize: file.FileSize,
				NodeID:   file.NodeID,
			}
			return file, nil
		}
	}
	return models.FileData{}, errors.New("Node not found")
}

func SelectNodeToCopyTo( fileNodes []string , dataNodes []models.DataNode) (string, error) {
	// alive node , not in the list of nodes that have the file
	validNodes := []string{}
	for _, node := range dataNodes {
		flag := false
		if node.IsDataNodeAlive == true {
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
		return "", errors.New("No valid nodes to copy to")
	} else {
		return validNodes[0], nil
	}

}

func CopyFileToNode(srcFile models.FileData, destNodeID string, dataNodes []models.DataNode) error {
	srcNodeID := srcFile.NodeID
	srcNode, err := GetNodeByID(srcNodeID,dataNodes)
	destNode, err := GetNodeByID(destNodeID,dataNodes)
	if err != nil {
		return fmt.Errorf("error getting node by ID: ", err)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", srcNode.IP,srcNode.Port), grpc.WithInsecure())
	if err != nil {
		
		return fmt.Errorf("error in dial: %v", err)
	}
	
	defer conn.Close()
	c := pb_r.NewDFSClient(conn)
	res, err := c.CopyNotification(context.Background(), &pb_r.CopyNotificationRequest{IsSrc: false, FileName: srcFile.Filename, FilePath: srcFile.FilePath, DestId: destNodeID, DestIp: destNode.IP, DestPort: destNode.Port})
	if err != nil {
		return fmt.Errorf("error in CopyNotification: ", err)
	}
	fmt.Println("CopyNotification response:", res.Ack)
	if res.Ack != "Ack" {
		return fmt.Errorf(res.Ack)
	}
	return nil
}
