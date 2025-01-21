package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func GetClassLoadRootPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	rootPath, err := filepath.Abs(filepath.Dir(exePath))
	if err != nil {
		return "", err
	}
	return rootPath, nil
}

func UploadUrl(fileName string) (string, string) {
	//path, err := GetClassLoadRootPath()
	staticDir := "./static/img"
	newName := fmt.Sprintf("/%s_%s", strconv.FormatInt(time.Now().UnixMilli(), 10), fileName)
	dst := staticDir + newName
	return dst, newName
}
