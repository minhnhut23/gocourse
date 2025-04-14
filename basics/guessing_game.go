package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random num between 1 & 100
	target := random.Intn(100) +1

	fmt.Println("Welcome to Guessing game!")
	fmt.Println("I have chosen a number between 1 & 100.")
	fmt.Println("Can u guess what it is?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulation! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too Low! Try guessing a higher	number.")
		} else if guess > target {
			fmt.Println("Too High! Try guessing a lower	number.")	
		}	
	}
}