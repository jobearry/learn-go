package main

import "fmt"

func main() {
	x := 22
	y := 9

	fmt.Println(isLess(x, y))

	fmt.Println(ifDemo(200))
	fmt.Println(switchDemo(200))
	fmt.Println(switchDemo(500))
	fmt.Println(switchDemo(301))
}

func isLess(x int, y int) bool {
	return x < y
}

func ifDemo(x int) string {
	if x == 200 {
		return "Success"
	} else {
		return "Error"
	}
}

func switchDemo(x int) string {
	switch {
	case x < 400:
		fmt.Println("x is less than 400")
		fallthrough
	case x == 200:
		return "Success"
	case x == 500:
		return "Server Error"
	default:
		return "Error"
	}
}
