package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string `json:"input_file_path"`
	OutputFilePath string `json:"output_file_path"`
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		errorMessage := "could not read file"
		return nil, errors.New(errorMessage)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		errorMessage := "reading the file failed"
		return nil, errors.New(errorMessage)
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to write data to file")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		inputPath,
		outputPath,
	}
}
