package main

import (
	"fmt"
	"log"
	"os"
)

//read directory and print info
func readCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	fileList, _ := file.Readdir(0)

	fmt.Printf("\nName\t\tSize\tIsDirectory  Last Modification\n")

	for _, files := range fileList {
		fmt.Printf("\nName: %-15s Size: %-7v Directory?: %-12v Modified: %v\n", files.Name(), files.Size(), files.IsDir(), files.ModTime())
	}
}

func main() {
	readCurrentDir()
}
