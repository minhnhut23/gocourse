package basics

import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Defferred i value:", i)
	defer fmt.Println("1st deffered statement executed")
	defer fmt.Println("2nd deffered statement executed")
	defer fmt.Println("3rd deffered statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println(i)
}