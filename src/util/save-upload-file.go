package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveUploadFile(originalFile *multipart.FileHeader) (fileName string, err error) {
	file, err := originalFile.Open()

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Create the uploads folder if it doesn't already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName = fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(originalFile.Filename))
	// create empty file
	dst, err := os.Create("./uploads/" + fileName)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
