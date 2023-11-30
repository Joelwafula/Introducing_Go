package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	//using open helps us get the contents of the directory.
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfo, err := os.ReadDir("1")

	if err != nil {
		return
	}

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}
}
