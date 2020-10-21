package main

import (
	"fmt"
)

func main() {
	var arr [10]int	//未初始化时，全部默认为0
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	sum := getSum(arr1)
	fmt.Println("the sum of arr1: ", sum)

	var arr2 [2][2]int
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[1]); j++ {
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}

	arr3 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[1]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}
}

func getSum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
