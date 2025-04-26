package intermediate

import (
	"fmt"
	"strings"
)

func main() {

	// str := "Hello Go!"

	// fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// // String conversion
	// num := 18
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // String splitting
	// fruit := "apple, orange, banana"
	// fruit1 := "apple-orange-banana"
	// parts := strings.Split(fruit, ",")
	// parts1 := strings.Split(fruit1, "-")
	// fmt.Println(parts)
	// fmt.Println(parts1)

	// countries := []string{"Germany", "France", "Italy"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go?"))

	// replaced := strings.Replace(str, "Go", "World", 1)
	// fmt.Println(replaced)

	// strwspace := " Hello Everyone "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.Repeat("foo ", 3))

	// fmt.Println(strings.Count("Hello", "l"))

	// fmt.Println(strings.HasPrefix("Hello", "He"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))

	// str5 := "Hello, 123 Go! 11"
	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString(str5, -1)
	// fmt.Println(matches)

	// str6 := "Hello, こんにちは"
	// fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// Convert builder to string
	result := builder.String()
	fmt.Println(result)

	// Using writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset Builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
