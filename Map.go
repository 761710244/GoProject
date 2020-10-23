package main

import (
	"fmt"
)

func main() {
	var hashMap map[string]string
	hashMap = make(map[string]string)
	hashMap["aa"] = "bb"
	hashMap["bb"] = "cc"

	for key, value := range hashMap {
		fmt.Println(key, value)
	}

	key, exist := hashMap["aa"]
	if exist {
		fmt.Println(key, hashMap[key])
	} else {
		fmt.Println("No")
	}

	//	initMap
	hashMap1 := map[int]int{1: 2, 2: 3, 3: 5, 6: 8}
	for key, value := range hashMap1 {
		fmt.Println(key, value)
	}
	//	delete
	delete(hashMap1, 2)
	for key, value := range hashMap1 {
		fmt.Println(key, value)
	}
}
