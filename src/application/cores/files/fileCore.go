package files

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

var RESOURCE_DIR string = "resource/"

func SetResourceDir(dir string) {
	if len(dir) == 0 || dir == "" {
		RESOURCE_DIR = GetCurrentDirectory() + RESOURCE_DIR
		return
	}
	RESOURCE_DIR = dir
}

func GetEnvPath() string {
	return GetCurrentDirectory() + "application.env"
}

func GetWebIndexPath() string {
	return GetCurrentDirectory() + "webs/index.html"
}

func GetAppJsPath() string {
	return GetCurrentDirectory() + "webs/app.js"
}

func GetResourceDirectory() string {
	return RESOURCE_DIR
}

func GetVideoEndedPath() string {
	return GetCurrentDirectory() + "webs/end.mp4"
}

func GetCurrentDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("[ERR] Get current directory: ", err)
	}
	return path + "/"
}

func CreateDirectory(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println("[ERR] Create directory: ", err)
		}
	}
}

func CreateFile(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("[ERR] Create file error", err)
	}
	defer f.Close()
}

func WriteFile(filePath, value string) {
	err := ioutil.WriteFile(filePath, []byte(value), 0644)
	if err != nil {
		log.Println("[ERR] Write : ", err)
	}
}

func ReadFile(filePath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("[ERR] Read : ", err)
	}
	return fileBytes, err
}

func ReadDirectory(directoryPath string) []fs.FileInfo {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func IsFileExit(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
