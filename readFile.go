package main

import (
	//"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		//handle the error
		return

	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

}
