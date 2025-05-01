package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	numString := "12345"
	num, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Int:", num)
	fmt.Println("Parsed Int:", num+1)

	numistr, err := strconv.ParseInt(numString, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Int:", numistr)

	floatStr := "3.14"
	floatValue, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatValue)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)

	invalidNum := "456abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed invalid to decimal:", invalidParse)

}
