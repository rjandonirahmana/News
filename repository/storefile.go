package repository

import (
	"errors"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

const AllowedExtensions = ".jpeg,.jpg,.png"

func StoreImage(imageName string, path string, file []byte) (string, error) {

	mime := mimetype.Detect(file)
	if strings.Index(AllowedExtensions, mime.Extension()) == -1 {
		return "", errors.New("File Type is not allowed, file type: " + mime.Extension())
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	folder, err := os.Create(path + imageName + mime.Extension())
	if err != nil {
		return "", err
	}
	defer folder.Close()

	_, err = folder.Write(file)
	if err != nil {
		return "", err
	}

	return folder.Name(), nil
}
