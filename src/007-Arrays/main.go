package main

import "fmt"

func main() {
	fmt.Println("Hello")

	arr1 := [2]int{2, 2}
	var arr2 [3]string

	fmt.Println(arr1)
	fmt.Println(arr2)

	for i, value := range arr1 {
		fmt.Println(i, value)
	}
}
