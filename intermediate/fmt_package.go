package intermediate

import "fmt"

func main() {

	// name := "a"
	// age := 12

	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	// fmt.Println(sf)

	// var name string
	// var age int

	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	err := checkAge(15)
	if err != nil {
		fmt.Println(err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young", age)
	}
	return nil
}
