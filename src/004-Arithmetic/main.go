package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := 8

	//
	fmt.Println(add(x, y))

	zz := x / y
	fmt.Println(zz)

	fname := "Jonathan"
	lname := "Golimlim"
	fmt.Println(concat(fname, lname))

	// num := "2209"
	result, err := stringToInt(fname)
	fmt.Println(result, err)
}

func add(x int, y int) int {
	return x + y
}

func concat(x string, y string) string {
	return x + y
}

func stringToInt(x string) (int, error) {
	return strconv.Atoi(x)
}
