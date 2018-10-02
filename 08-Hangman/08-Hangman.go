// 08: Hangman
package main

import (
	"fmt"
	"math/rand"
)

type hangman struct {
	word     string
	guesses  string
	tries    int
	maxTries int
}

func main() {
	// Select a word from the dict
	word := dict[rand.Int()%len(dict)]

	game := hangman{word: word, maxTries: 5}

	// Display welcome and run game loop
	fmt.Println("08: Hangman")
	fmt.Printf("A word has been chosen! It is %v letters long.", len(word))
	fmt.Printf("You have 5 strikes before Sussman gets noos'd")

	var strikes int
	for {
		game.displayBoard()

	}
}

func (h *hangman) displayBoard() {

}
func (h *hangman) getGuess() {
}
