package utils

import (
	"os"
	"strings"
	"fmt"
	"bufio"
)


func ReadFile(masterAddress,clientAddress *string,nodes *[]string)(){
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
		switch key {
		case "Master_Address":
			*masterAddress = value
		case "Client_Address":
			*clientAddress = value
		default:
			if strings.HasPrefix(key, "Node") {
				*nodes = append(*nodes,value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}
}
}
