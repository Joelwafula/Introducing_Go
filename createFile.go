package main

//here we are creating a file, using the os package
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")

	if err != nil {
		//an  error execuable statement will be returned
		fmt.Println("Opps the above working file crashed")
		return
	}
	defer file.Close()

	file.WriteString("testing this file if its working")
}
