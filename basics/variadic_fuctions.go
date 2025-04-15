package basics

import "fmt"

func main() {
	// ... Ellipsis
	// func funcName(param1 type1, param3 ...type3) returnType{
	// 	function body
	// }

	// fmt.Println(sum(1, 2, 3))

	// statement, total := sum("Sum of 1, 2, 3 is", 1, 2, 3)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence, total := sum(3, numbers...)
	fmt.Println(sequence, total)


}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}