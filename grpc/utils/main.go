package utils

import (
	"os"
	"strings"
	"fmt"
	"bufio"
)

type Node struct {
	IP   string
	Port string
}

func ReadFile(masterAddress,clientAddress *string,nodes *[]Node)(){
	// Open the file
	file, err := os.Open("config.txt")
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
		valueParts := strings.Split(value, ":")
		ip := valueParts[0]
		port := valueParts[1]

		switch key {
		case "Master_Address":
			*masterAddress = value
		case "Client_Address":
			*clientAddress = value
		default:
			if strings.HasPrefix(key, "Node") {
				*nodes = append(*nodes, struct {
					IP   string
					Port string
				}{
					IP:   ip,
					Port: port,
				})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}
}
}
