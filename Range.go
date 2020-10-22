package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7}
	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)

	//	index为下标，num为数组的值
	for index, num := range arr {
		fmt.Println(index, num)
	}

	//	ASCII码
	for index, val := range "abc" {
		fmt.Println(index, val)
	}
}
