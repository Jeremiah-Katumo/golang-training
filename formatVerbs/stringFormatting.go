package main

import (
	"fmt"
)

func main() {
	var txt = "Hello"

	fmt.Printf("%s\n", txt)   	// %s Prints the value as plain string
	fmt.Printf("%q\n", txt)   	// %q Prints the value as a double-quoted string
	fmt.Printf("%8s\n", txt)  	// %8s Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", txt) 	// %-8s Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", txt)		// %x Prints the value as hex dump of byte values
	fmt.Printf("% x\n", txt)	// % x Prints the value as hex dump with spaces
}
