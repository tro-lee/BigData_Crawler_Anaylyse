package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func GetFlagPath() (string, error) {
	path := flag.String("path", "", "path to file")
	flag.Parse()

	if *path == "" {
		return "", fmt.Errorf("path is required")
	}
	return *path, nil
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func JsonToFile(jsonData []byte, filePath string) error {
	file, err := os.Create(filePath)
	file.Write(jsonData)
	if err != nil {
		err = os.Mkdir("./result", 0777)
		if err != nil {
			return err
		}

		JsonToFile(jsonData, filePath)
	}
	return nil
}
