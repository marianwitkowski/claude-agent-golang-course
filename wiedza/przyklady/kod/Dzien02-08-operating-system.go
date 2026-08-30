package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createFile() {
	fd, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	fd.Close()
}

func createDir() {
	fileInfo, err := os.Stat("test")
	if err == nil {
		fmt.Println(fileInfo.Name())
		fmt.Println(fileInfo.Size())
		fmt.Println(fileInfo.ModTime())
		fmt.Println(fileInfo.IsDir())
		fmt.Println(fileInfo.Mode())
	}
	err = os.MkdirAll("test", 0666)
	if err != nil {
		println(err)
	}
}

func osUtils() {
	fmt.Println(os.Environ())
	fmt.Println(os.Getpid())
	fmt.Println(os.Getwd())
	os.Rename("app.log", "debug.log")
}

func findFiles() {
	root := "/Users/marian/golang20230414/Dzien01"
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		return nil
	})
}

func findFilesWithMask() {
	root := "/Users/marian/golang20230414/Dzien01"
	mask := "*.txt"
	files, _ := filepath.Glob(filepath.Join(root, mask))
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	createDir()
	osUtils()
	findFilesWithMask()
}
