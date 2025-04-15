package basics

import "fmt"

func main() {
	mes := "Hello World!"
	for i, v := range mes {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}