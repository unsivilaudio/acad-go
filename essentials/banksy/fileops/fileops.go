package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FileIO struct {
	fileName string
}

func CreateFileStream(fileName string) *FileIO {
	stream := FileIO{
		fileName: fileName,
	}
	return &stream
}

func (f *FileIO) WriteFloatToFile(num float64) {
	valueText := fmt.Sprintf("%.2f", num)
	os.WriteFile(f.fileName, []byte(valueText), 0644)
}

func (f *FileIO) GetFloatFromFile(initial float64) (float64, error) {
	fileValue, err := os.ReadFile(f.fileName)
	if err != nil {
		errorText := fmt.Sprintf("Failed to find %v file.", f.fileName)
		return initial, errors.New(errorText)
	}
	valueText := string(fileValue)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		errorText := fmt.Sprintf("Failed to parse %v value.", f.fileName)
		return initial, errors.New(errorText)
	}
	return value, nil
}
