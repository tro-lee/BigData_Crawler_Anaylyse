package utils

import (
	"os"
)

func JsonToFile(jsonData []byte, filePath string) error {
	file, err := os.Create(filePath)
	file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
