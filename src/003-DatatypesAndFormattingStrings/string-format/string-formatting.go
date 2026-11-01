package main

import "fmt"

func main() {
	x := 7
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
	fmt.Printf("%b\n", x)

	y := 3.14159

	fmt.Printf("%e\n", y)
	fmt.Printf("%f\n", y)

	z := "JG"
	fmt.Printf("%s\n", z)

	//format precisions
	fmt.Printf("%.2f\n", y)
	fmt.Printf("%10.2f\n", y)
	fmt.Printf("%.2f%%\n", y)

	//format a string then store to a variable
	xx := fmt.Sprintf("%10.2f", y)
	fmt.Println(xx, xx, xx)
}
