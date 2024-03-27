package main

import (
	"fmt"
)

func main() {
	var i = 3.141

	fmt.Printf("%e\n", i) // %e	Scientific notation with 'e' as exponent
	fmt.Printf("%f\n", i) // %f	Decimal point, no exponent
	fmt.Printf("%.2f\n", i) // %.2f	Default width, precision 2
	fmt.Printf("%6.2f\n", i) // %6.2f	Width 6, precision 2
	fmt.Printf("%g\n", i) // %g	Exponent as needed, only necessary digits
}
