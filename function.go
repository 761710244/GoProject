package main

import (
	"fmt"
	"math"
)

type Tangle struct {
	length int
	width  int
}

func main() {
	var a, b int = 8, 6
	fmt.Println(getMax(a, b))

	var c, d int = 10, 20
	fmt.Println(max(c, d))

	var e, f string = "abc", "def"
	fmt.Println(swap(e, f))

	var x, y int = 100, 200

	fmt.Println("before value Go:")
	fmt.Println(x, y)

	fmt.Print("after value Go:")
	valueGo(x, y)
	fmt.Println(x, y)

	fmt.Print("after another Go:")
	anotherGo(&x, &y)
	fmt.Println(x, y)

	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSqrt(100))

	var O1 Tangle
	O1.length = 10
	O1.width = 8
	fmt.Println("Area is: ", O1.getArea())
}

func (o Tangle) getArea() int {
	return o.width * o.length
}

func getMax(num1, num2 int) int {
	var res int
	if num1 >= num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

/**
 *	值传递
 */
func valueGo(x, y int) int {
	var tmp int
	tmp = x
	x = y
	y = tmp
	return tmp
}

/**
 *	引用传递
 */
func anotherGo(x *int, y *int) int {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
	return tmp
}
