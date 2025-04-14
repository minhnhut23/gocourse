package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	// Overflow with signed int
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned int
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point number
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}