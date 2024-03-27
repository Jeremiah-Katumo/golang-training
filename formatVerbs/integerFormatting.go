package main
import ("fmt")

func main() {
  var i = 15
 
  fmt.Printf("%b\n", i)     // Base 2
  fmt.Printf("%d\n", i)		// Base 10
  fmt.Printf("%+d\n", i)    // Base 10 and always show sign
  fmt.Printf("%o\n", i)		// Base 8
  fmt.Printf("%O\n", i)		// Base 8, with leading 0o
  fmt.Printf("%x\n", i)		// Base 16, lowercase
  fmt.Printf("%X\n", i)		// Base 16, uppercase
  fmt.Printf("%#x\n", i)	// Base 16, with leading 0x
  fmt.Printf("%4d\n", i)	// Pad with spaces (width 4, right justified)
  fmt.Printf("%-4d\n", i)	// Pad with spaces (width 4, left justified)
  fmt.Printf("%04d\n", i)	// Pad with zeroes (width 4)
}