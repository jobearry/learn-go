package main

import "fmt"

func main() {
	countToTen()

	str := "Jonathan"
	runThroughString(str)
	fmt.Print("\n")
}

func countToTen() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
}

func runThroughString(x string) {
	for _, char := range x {
		fmt.Printf("%c", char)
	}
}
