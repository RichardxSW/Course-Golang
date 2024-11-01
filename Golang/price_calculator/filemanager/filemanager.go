package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"time"

	// "fmt"
	"os"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error){
	file, err := os.Open(fm.InputFilePath)

	if err != nil { 
		return nil, errors.New("Could not open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Failed to read line in file")
	}

	// file.Close()
	return lines, nil
}

// interface bisa diganti any
func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	
	if err != nil {
		return errors.New("Failed to create file")
	}

	defer file.Close()

	time.Sleep(time.Second * 2)

	encoder :=json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	// file.Close()
	return nil
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}