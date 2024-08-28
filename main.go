package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guesser Game!")
	fmt.Println("A random number will be sorted. Try to guess a number between 0 and 100.")

	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}
	x := rand.Int64N(101)

	for i := range guesses {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Invalid guess. Please enter an integer number.")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("Too low. Try again.")
		case guessInt > x:
			fmt.Println("Too high. Try again.")
		case guessInt == x:
			fmt.Println("Correct! You guessed it in", i+1, "attempts.")
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"Sorry, you ran out of guesses."+
			" The number was %d.\n"+
			"This was your guesses: %v\n", x, guesses)
}
