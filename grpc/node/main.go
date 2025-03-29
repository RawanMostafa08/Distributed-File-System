package main

import (
	"context"
	"fmt"
	"net"
	"io/ioutil"
	"os"
	"io"
	"time"
	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	pbHeartBeats "github.com/RawanMostafa08/Distributed-File-System/grpc/HeartBeats"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

type textServer struct {
	pb.UnimplementedDFSServer
}

func (s *textServer) DownloadFileRequest(ctx context.Context, req *pb.DownloadFileRequestBody) (*pb.DownloadFileResponseBody, error) {
	fmt.Println("DownloadFileRequest called for:", req.FileName, "Range:", req.Start, "-", req.End)

	// Open the file
	filePath := "grpc\\files\\" + req.FileName
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

// ReadMP4File reads an MP4 file and returns its content as a byte slice.
func ReadMP4File(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
// func pingMaster(node_index int32,masterAddress string) ( error) {
// 	heartBeat := &pbHeartBeats.HeartbeatRequest{NodeId: "Node_" + fmt.Sprint(node_index)}
// 	for {
// 		print("Node_" + fmt.Sprint(node_index) + " is pinging master...")
// 		conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
// 		if err != nil {
// 			fmt.Println("did not connect:", err)
// 			return  err
// 		}
// 		master := pbHeartBeats.NewHeartbeatServiceClient(conn)
// 		fmt.Println("Connected to Master",conn)
// 		master.KeepAlive(context.Background(), heartBeat)
// 		if err != nil {
// 			log.Fatal("Error in pinging master:", err)
// 			return err
// 		}
// 		time.Sleep(time.Second)
// 	}
// }
func pingMaster(nodeIndex int32, masterAddress string) error {
    heartBeat := &pbHeartBeats.HeartbeatRequest{NodeId: "Node_" + fmt.Sprint(nodeIndex)}

    // Move connection setup outside loop
    conn, err := grpc.Dial(masterAddress, grpc.WithInsecure())
    if err != nil {
        fmt.Println("Did not connect:", err)
        return err
    }
    defer conn.Close()

    master := pbHeartBeats.NewHeartbeatServiceClient(conn)

    for {
        fmt.Println("Node_" + fmt.Sprint(nodeIndex) + " is pinging master...")

        // // Use a timeout to prevent blocking
        // ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        // defer cancel()

        _, err := master.KeepAlive(context.Background(), heartBeat)
        if err != nil {
            fmt.Println("Error in pinging master:", err)
        } else {
            fmt.Println("Ping successful")
        }

        time.Sleep(time.Second)
    }
}




func main() {

	// Read input from user
	fmt.Print("Enter node index : ")
	var node_index int32
	fmt.Scanln(&node_index)
		

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	
	lis, err := net.Listen("tcp", nodes[node_index])
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterDFSServer(s, &textServer{})
	fmt.Println("Server started. Listening on port ",nodes[node_index],"...")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve:", err)
	}
	go pingMaster(node_index,masterAddress)
}
