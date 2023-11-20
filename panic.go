package main

import "fmt"

// func main() {
// 	panic("PANIC")

// 	str := recover() ->THIS ONE

// 	fmt.Println(str)

// }

// this is not a correct one above, because panic immediately causes the function to stop execution

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)

	}()
	panic("PANIC")
}

//In the above code, recover should be implemented with defer
