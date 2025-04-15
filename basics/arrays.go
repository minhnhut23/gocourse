package basics

import "fmt"

func main() {

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"apple", "banana", "orange", "grapes"}
	// fmt.Println(fruits)

	// fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	var copyArray *[3]int

	copyArray = &originalArray
	copyArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", *copyArray)

	// for i:=0; i<len(numbers); i++{
	// 	fmt.Println("Element at index,", i, ":", numbers[i])
	// }

	// for _, value := range numbers {
	// 	fmt.Printf("Value: %d\n", value)
	// }

	// fmt.Println("The length of numbers array is", len(numbers))

	// // Comparing Arrays
	// array1 := [3]int{1,2,3}
	// array2 := [3]int{10,2,3}

	// fmt.Println("Array1 is equal to Array2:", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }

	// fmt.Println(matrix)
		
}