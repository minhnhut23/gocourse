package basics

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(param1 type1, param2 type2,...) (returnType1, returnType2) {
	// code block
	// return value1, value2
	// }

	q, r := divide(10, 3)
	fmt.Println(q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error){
	if (a > b){
		return "a > b", nil
	} else if (a < b){
		return "a < b", nil
	}
	return "", errors.New("Unable to compare which is greaterz")
}