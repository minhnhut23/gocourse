package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Structs, interface, enums

	// snake_case

	// UPPERCASE
	// Use case in Constants

	// mixedCase

	const MAXRETIRED = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}