package util

import (
	"fmt"
	"io"
	"os"
)

func GetInput(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("could not read file: %v", err)
		}
	}
	return b, nil
}
