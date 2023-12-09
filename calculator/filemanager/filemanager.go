package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

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
