package main

import (
	// "context"
	"fmt"
	// "net"
	// "strings"

	pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/peer"
)

// type DataNode struct {
// 	DataKeeperNode string
// 	IsDataNodeAlive bool
// }

// type FileData struct {
// 	Filename string
// 	FilePath string
// }

// type LookUpTableTuple struct {
// 	Node DataNode
// 	File []FileData
// }

// type textServer struct {
// 	pb.UnimplementedTextServiceServer
// }

// func (s *textServer) Capitalize(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
// 	text := req.GetText()
// 	p, _ := peer.FromContext(ctx)
// 	fmt.Println("Request from:", p.Addr)
// 	capitalizedText := strings.ToUpper(text)
// 	return &pb.TextResponse{CapitalizedText: capitalizedText}, nil
// }

func main() {
	fmt.Println("master started...")
	pb.UploadRequestBody()
	// lis, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	fmt.Println("failed to listen:", err)
	// 	return
	// }
	// s := grpc.NewServer()
	// pb.RegisterTextServiceServer(s, &textServer{})
	// fmt.Println("Server started. Listening on port 8080...")
	// if err := s.Serve(lis); err != nil {
	// 	fmt.Println("failed to serve:", err)
	// }
}
