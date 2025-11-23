package main

import (
	"os"
	"fmt"
)

func main() {	
	if len(os.Args) < 2 {
		fmt.Println("Please provide input file. Example: go run main.go commands.txt")
		return
	}

	filePath := os.Args[1]

	var fileHandler *FileHandler
	fileHandler = NewFileHandler(filePath)
	fileHandler.Execute()	
	fileHandler.Close()

	os.Exit(0)
}
