package main

import (
	// "context"
	// "fmt"
	pbUtils "github.com/RawanMostafa08/Distributed-File-System/grpc/utils"
	// pb "github.com/RawanMostafa08/Distributed-File-System/grpc/Upload"// Import the generated package
	// "google.golang.org/grpc"
)

func main() {

	var masterAddress, clientAddress string
	nodes := []string{}

	pbUtils.ReadFile(&masterAddress,&clientAddress,&nodes)
	// fmt.Println("Master Address:", masterAddress)
	// fmt.Println("Client Address:", clientAddress)
	// fmt.Println("Nodes Master:", nodes)
}
