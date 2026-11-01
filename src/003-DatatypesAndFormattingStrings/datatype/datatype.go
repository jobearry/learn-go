package main

import (
	"fmt"
	"strconv"
)

func main() {
	// uint
	// int
	// float
	// rune = int32
	// byte = int8
	// bool
	// string
	// nil

	//=== explicit assignment ===//
	var test int = 0
	test += 1

	const test1 string = "Hello"

	fmt.Println(test1 + " " + strconv.Itoa(test))

	//=== implicit assignment ===//
	test2 := "__JG__"
	fmt.Printf("%T\n", test2)
}
