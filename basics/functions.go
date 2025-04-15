package basics

import "fmt"

func main() {
	//func <name>(parameter list) returnType{
	// code block
	// return value
	// }

	// sum := add(1, 2)
	// fmt.Println(add(1, 2))

	// greet := func(){
	// 	fmt.Println("Hello Annoymous Fuction")
	// }

	// greet()

	// operation := add

	// result := operation(3, 5)
	// fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println(result)

	// Returning and using a fuction
	multiplBy2 := createMultiplier(2)
	fmt.Println(multiplBy2(6))
}

func add(a, b int) int {
	return a + b
}

// Function that take a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that return a fuction
func createMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}