package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, Go"
	message1 := "Hello, \tGo"
	message2 := "Hello, \rGo" // Gollo
	rawMessage := "Hello\nGo"

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Len of message:", len(message))
	fmt.Println("The first of message:", message[0])

	greeting := "Hello "
	name := "John"
	fmt.Println(greeting + name)

	str := "Apple"   // A has an ASCII value of 65
	str1 := "apple"  // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97

	fmt.Println(str < str2)
	fmt.Println(str3 < str)
	fmt.Println(str1 > str)
	fmt.Println(str1 > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr %T\n", cstr)

	const NIHONGO = "海洋"
	fmt.Println(NIHONGO)

	jHello := "こんにちは"
	for _, runeValue := range jHello {
		fmt.Printf("%v\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
