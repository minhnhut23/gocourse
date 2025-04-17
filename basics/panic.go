package basics

import "fmt"

func main() {

	// panic(interface())

	// Example of a valid input
	process(10)

	// Example of an invalid input
	process(-10)

}

func process(input int) {
	defer fmt.Println("Defferd 1")
	defer fmt.Println("Defferd 2")
	
	if input < 0 {
		fmt.Println("Before panic")
		panic("Input must be non-negative number")
	}
	fmt.Println("Processing input:", input)
}