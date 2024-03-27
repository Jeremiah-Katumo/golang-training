package main
import ("fmt")

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// To only show the value or the index, you can omit the other output using an underscore (_).
	myfruits := [3]string{"apple", "orange", "banana"}
	for _, val := range myfruits {
		fmt.Printf("%v\n", val)
	}
}