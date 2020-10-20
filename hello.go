package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	var a = "1234"
	fmt.Println(a)

	aaa := 122
	fmt.Println(aaa)

	var b, c int = 6, 8
	fmt.Println(b, c)

	/**
	未初始化的整型是0
	*/
	var d int
	fmt.Println(d)

	/**
	未初始化的布尔型为false
	*/
	var e bool
	fmt.Println(e)

	const lenth int = 10
	const windth int = 8
	var area int
	area = lenth * windth
	fmt.Println(area)

	const (
		aa = "abc"
		bb = "def"
		cc = "ghi"
	)
	fmt.Println(aa, bb, cc)
}
