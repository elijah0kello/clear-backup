package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// instantiate the directory path and read the directory contents
	dirName := "/Users/macbook/Desktop/go-learn/"
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// record the current time
	now := time.Now()

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Println(err)
			return
		}
		oneWeek := 7 * 24 * time.Hour
		age := now.Sub(fileInfo.ModTime())
		if age > oneWeek {
			fmt.Println("Deleting file ...", file.Name())
			err := os.Remove(dirName + file.Name())
			if err != nil {
				fmt.Println("Deletion failed with error:", err)
				return
			}
			fmt.Println("File has been deleted")
		} else {
			fmt.Println("File will not be deleted")
		}
	}

}
