package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane" //type is inferred
	x := 2 //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// variable declaration without initial value
	var am string
	var bm int
	var cm bool

	fmt.Println(am)
	fmt.Println(bm)
	fmt.Println(cm)

	// value assignment after declaration
	var student3 string
	student3 = "John"
	fmt.Println(student3)

	// multiple variable declaration
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// Variable declaration in a block
	var (
		ad int
		bd int = 1
		cd string = "hello"
	)
   
	fmt.Println(ad)
	fmt.Println(bd)
	fmt.Println(cd)
}