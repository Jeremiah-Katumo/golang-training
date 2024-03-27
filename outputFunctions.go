package main

import (
	"fmt"
)

func main() {
	var i,j string = "Hello","World"

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// Println()
	var a, b string = "Hello", "World"

	fmt.Println(a, b)

	// %v is used to print the value of the arguments
	// %T is used to print the type of the arguments
	var c string = "Hello"
	var d int = 15

	fmt.Printf("i has value: %v and type: %T\n", c, c)
	fmt.Printf("j has value: %v and type: %T", d, d)
}