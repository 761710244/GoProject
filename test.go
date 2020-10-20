package main

import "fmt"

func main() {

	var a int = 6
	if a == 6 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	var str string = "abc"
	var res int = 0
	switch str {
	case "aa":
		res = 1
	case "abb":
		res = 2
	case "abc", "bbc":
		res = 3
	default:
		res = -1
	}
	fmt.Println(res)
}
