package Replicate_utils

import (
	"context"
	"errors"
	"fmt"

	pb_r "github.com/RawanMostafa08/Distributed-File-System/grpc/Replicate"
	"github.com/RawanMostafa08/Distributed-File-System/grpc/models"
	"google.golang.org/grpc"
)

// Replicate Helper Functions

func GetNodeByID(nodeID string, dataNodes []models.DataNode) (models.DataNode, error) {
	for _, node := range dataNodes {
		if node.NodeID == nodeID {
			return node, nil
		}
	}
	return models.DataNode{}, errors.New("node not found")
}

func GetFileNodes(fileName string, lookupTable []models.FileData, dataNodes []models.DataNode) []string {
	nodes := []string{}
	for _, f := range lookupTable {
		filenode, err := GetNodeByID(f.NodeID, dataNodes)
		if err != nil {
			fmt.Println("Error getting node by ID:", err)
		} else if f.Filename == fileName && filenode.IsDataNodeAlive {
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
	return models.FileData{}, errors.New("node not found")
}

func SelectNodeToCopyTo(fileNodes []string, dataNodes []models.DataNode) (string, string, error) {
	// alive node , not in the list of nodes that have the file

	for _, node := range dataNodes {
		flag := false
		if node.IsDataNodeAlive {
			for _, fileNode := range fileNodes {
				if fileNode == node.NodeID {
					flag = true
					break
				}
			}
			if !flag {
				for i, port := range node.Port {
					if !node.IsPortBusy[i] {
						return node.NodeID, port, nil
					}
				}
			}
		}
	}
	return "", "", errors.New("no available nodes to copy to")

}

func CopyFileToNode(srcFile models.FileData, destNodeID string, destNodePort string, dataNodes *[]models.DataNode) error {
	srcNodeID := srcFile.NodeID
	srcNode, err := GetNodeByID(srcNodeID, *dataNodes)
	if err != nil {
		return fmt.Errorf("error getting src node by id: %v", err)
	}
	destNode, err := GetNodeByID(destNodeID, *dataNodes)
	if err != nil {
		return fmt.Errorf("error getting dest node by id: %v", err)
	}
	/////////////////////
	srcPort := ""
	for i, port := range srcNode.Port {
		if !srcNode.IsPortBusy[i] {
			srcPort = port
		}
	}
	for _, node := range *dataNodes {
		// src port busy
		if node.NodeID == srcNodeID {
			for j, port := range node.Port {
				if port == srcPort {
					node.IsPortBusy[j] = true
				}
			}
			//dest port busy
		} else if node.NodeID == destNodeID {
			for j, port := range node.Port {
				if port == destNodePort {
					node.IsPortBusy[j] = true
				}
			}
		}
	}

	if srcPort == "" {
		return fmt.Errorf("no free port found on source node")
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", srcNode.IP, srcPort), grpc.WithInsecure(), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*200), // 200MB receive
		grpc.MaxCallSendMsgSize(1024*1024*200), // 200MB send
	))
	if err != nil {
		return fmt.Errorf("error in dial: %v", err)
	}

	defer conn.Close()
	c := pb_r.NewDFSClient(conn)
	res, err := c.CopyNotification(context.Background(), &pb_r.CopyNotificationRequest{IsSrc: false, FileName: srcFile.Filename, FilePath: srcFile.FilePath, DestId: destNodeID, DestIp: destNode.IP, DestPort: destNodePort})
	if err != nil {
		return fmt.Errorf("error in CopyNotification: %v", err)
	}
	fmt.Println("CopyNotification response:", res.Ack)
	if res.Ack != "Ack" {
		return fmt.Errorf("%s", res.Ack)
	}
	for _, node := range *dataNodes {
		// src port free
		if node.NodeID == srcNodeID {
			for j, port := range node.Port {
				if port == srcPort {
					node.IsPortBusy[j] = false
				}
			}
			//dest port free
		} else if node.NodeID == destNodeID {
			for j, port := range node.Port {
				if port == destNodePort {
					node.IsPortBusy[j] = false
				}
			}
		}
	}
	return nil
}
