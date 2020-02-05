package main

import (
	"fmt"
)

const favColor string = "blue"

func main() {
	var guess string
	// Create an input loop
	for {
		// Ask the user to guess the favorite color
		fmt.Println("Guess my favorite color...")
		// Try to read a line of input from user. Print out the error 0
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		// Did they guess the correct answer
		if favColor == guess {
			// They guessed right!
			fmt.Printf("%q is my favorite color!\n", favColor)
			return
		}
		// Wrong! Have the at the guess loop
		fmt.Printf("Sorry, %q is not my favorite color. Guess again!\n", guess)
	}
}
