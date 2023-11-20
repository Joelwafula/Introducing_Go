package main

import "fmt"

func main() {
	x := map[string]map[string]string{
		"gas1": map[string]string{
			"helium": "used to burn",
			"state":  "gaseous",
		},
		"gas2": map[string]string{
			"argon": "inate",
			"state": "stateless",
		},
	}
	fmt.Println(x)
}
