package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type FileHandler struct {
	filePath string
	file	 *os.File
}

func NewFileHandler(path string) *FileHandler {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return &FileHandler{
		filePath: path,
		file:     file,
	}
}

func (fh *FileHandler) Execute(){
	var parking *ParkingLot

	scanner := bufio.NewScanner(fh.file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {

		case "create_parking_lot":			
			n, _ := strconv.Atoi(parts[1])
			parking = NewParkingLot(n)

		case "park":
			parking.Park(parts[1])

		case "leave":
			hours, _ := strconv.Atoi(parts[2])
			parking.Leave(parts[1], hours)

		case "status":
			parking.Status()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func (fh *FileHandler) Close(){
	fh.file.Close()
	return
}