package intermediate

import "fmt"

func main() {

	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	// fmt.Println(ptr)
	// fmt.Println(*ptr)
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }
	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()

	if ptr == nil {
		panic("Pointer is nil")
	}

	*ptr++
}
