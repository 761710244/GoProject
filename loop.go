package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum <= 60 {
		sum += sum
	}
	fmt.Println(sum)

	//for {
	//	fmt.Println("123")
	//}

	var a, b int
	for a = 1; a <= 10; a++ {
		if a == 6 {
			fmt.Println(a)
			break
		}
	}
	for b = 1; b <= 10; b++ {
		if b == 6 {
			fmt.Println(a - 1)
			continue
		}
	}
	fmt.Println("over")
}
