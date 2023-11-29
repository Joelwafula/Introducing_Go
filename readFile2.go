package main

import (
	"fmt"
	//"io/ioutil"
	"os"
)

func main() {

	bs, err := os.ReadFile("testdata/hello")

	if err != nil {
		return
	}
	//os.Stdout.Write(data)
	str := string(bs)
	fmt.Println(str)

}
