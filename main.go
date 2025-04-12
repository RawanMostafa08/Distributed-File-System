package utils

import (
	"os"
	"strings"
	"fmt"
	"bufio"
)


func ReadFile_master(masterAddress *string,nodes *[]string)(){
	// Open the file
	file, err := os.Open("../config_master.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "Master_Address":
			*masterAddress = value
		case "Node0_Address":
			*nodes = append(*nodes, value)
		case "Node1_Address":
			*nodes = append(*nodes, value)
		case "Node2_Address":
			*nodes = append(*nodes, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}
}

func ReadFile_node(masterAddress *string,nodes *[]string)(){
	// Open the file
	file, err := os.Open("../config_node.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "Master_Address":
			*masterAddress = value
		case "Node0_Address":
			*nodes = append(*nodes, value)
		case "Node1_Address":
			*nodes = append(*nodes, value)
		case "Node2_Address":
			*nodes = append(*nodes, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}
}

func ReadFile_client(masterAddress , clientAddress *string)(){
	// Open the file
	file, err := os.Open("../config_client.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "Master_Address":
			*masterAddress = value
		case "Client_Address":
			*clientAddress = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}
}