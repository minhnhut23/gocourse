package intermediate

import "fmt"

func main() {
	num := 0
	fmt.Print("Enter numbers of fibonacies: ")
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {

	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
