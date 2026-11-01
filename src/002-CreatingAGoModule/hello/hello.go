package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("Jobry")
	fmt.Println(message)
}
