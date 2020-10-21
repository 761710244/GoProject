package main

import "fmt"

func main() {
	var arr1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//arr2 := []int{1, 2, 3}
	//arr3 := make([]int, 1, 2)	//len cap,其中1为长度，2为容量

	//	获取长度
	fmt.Println(len(arr1))
	//	获取容量
	fmt.Println(cap(arr1))

	printSlice(arr1)
	//	截取部分数据
	printSlice(arr1[1:6])
	arr1 = append(arr1, 10)
	printSlice(arr1)
}

func printSlice(arr []int) {
	fmt.Print(arr)
	fmt.Println()
}
