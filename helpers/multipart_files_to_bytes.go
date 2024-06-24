package helpers

import (
	"io"
	"log"
	"mime/multipart"
)

func FileToByteArray(file *multipart.FileHeader) ([]byte, error) {
	openedFile, _ := file.Open()

	binaryFile, err := io.ReadAll(openedFile)

	if err != nil {
		return nil, err
	}

	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			log.Fatalf("Failed closing file %v", file.Filename)
		}
	}(openedFile)

	return binaryFile, nil
}
