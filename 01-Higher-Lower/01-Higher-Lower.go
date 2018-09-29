//01: Higher/Lower, Heads/Tails
// I guess we play some games?

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

var stdin *bufio.Reader

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var game string

	// Display menu
	fmt.Println("01: Higher/Lower, Heads/Tails")
	fmt.Println(`
Would you like to play a game:
(L) Higher or Lower
(T) Heads or Tails
(E) Exit`)

	// Select Game
	for {
		fmt.Scanln(&game)
		if game == "T" || game == "L" || game == "Global Thermonuclear War" {
			break
		}
		if game == "E" {
			return
		}
		fmt.Println("Uhhhhh. Let's try that again.")
	}

	//Transfer execution to subroutine
	if game == "L" {
		playHighLow()
	} else if game == "T" {
		playHeadsTails()
	} else if game == "Global Thermonuclear War" {
		fmt.Println("How about a nice game of chess instead Dr?")
		return
	} else {
		fmt.Println("I'm afarid I can't do that.")
	}
	fmt.Println("Goodbye.")
}

func playHeadsTails() {
	totalWins := 0
	var in string

	fmt.Println("We will play 5 games!")
	for i := 0; i < 5; i++ {
		fmt.Println("(H) Heads or (T) Tails?")
		fmt.Scanln(&in)

		// Heads is 0, Tails is 1
		coin := rand.Int() % 2
		if coin == 0 && in == "H" {
			fmt.Println("You guessed correctly!")
			totalWins++
		} else if coin == 1 && in == "T" {
			totalWins++
			fmt.Println("You guesed correctly!")
		} else {
			fmt.Println("That was incorrect!")
		}
	}

	fmt.Printf("You won %v of 5 games!\n", totalWins)
}
func playHighLow() {
	totalWins := 0
	var in string
	fmt.Println("All numbers will be between 0-100\nWe will play 5 games!")

	for i := 0; i < 5; i++ {
		compNumber := rand.Int() % 100
		fmt.Printf("The number is %v\n(H)igher or (L)ower?\n", compNumber)
		fmt.Scanln(&in)

		// Verify input
		if !(in == "H" || in == "L") {
			fmt.Println("You can't even type! That's a loss")
		}

		// Generate and check for win
		// Computer takes the win in a tie or invalid
		answer := rand.Int() % 100
		if answer < compNumber && in == "L" {
			fmt.Println("You guesed correctly!")
			totalWins++
		} else if answer > compNumber && in == "H" {
			fmt.Println("You guesed correctly!")
			totalWins++
		} else {
			fmt.Println("You tried buddy!")
		}
	}

	fmt.Printf("You won %v of 5 games!\n", totalWins)
}
