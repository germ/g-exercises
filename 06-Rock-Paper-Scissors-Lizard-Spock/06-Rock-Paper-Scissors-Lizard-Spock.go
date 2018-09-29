// 03: Age In Seconds (Including Leap)
package main

import (
	"fmt"
)

// Identifiers and names of choices
const (
	Rock     = 1 << iota
	Paper    = 1 << iota
	Scissors = 1 << iota
	Lizard   = 1 << iota
	Spock    = 1 << iota
)

// Necessary because the above are just ints
var names = map[int]string{
	Rock:     "Rock",
	Paper:    "Paper",
	Scissors: "Scissors",
	Lizard:   "Lizard",
	Spock:    "Spock",
}

// Choices for combat, when Anded together it yeilds
// exactly one indetifier! (Except in tie)
var loseTable = map[int]int{
	Rock:     Rock | Paper | Spock,
	Paper:    Paper | Lizard | Scissors,
	Scissors: Scissors | Spock | Rock,
	Lizard:   Lizard | Scissors | Rock,
	Spock:    Spock | Paper | Lizard,
}

// I'm not going to explain how RPSLS works
// We represent each possible play int, retrieving
// the winners name from a table. Manually handling
// the tie scenario
func main() {
	fmt.Println("06: Rock Paper Scissors Lizard Spock")
	fmt.Println("We will play 3 games.")
	for i := 0; i < 3; i++ {
		p1, p2 := getChoice("Player 1"), getChoice("Player 2")
		winner := loseTable[p1] & loseTable[p2]

		if p1 == p2 {
			fmt.Printf("It's a tie!\n\n")
		} else {
			fmt.Printf("The winner is %v\n\n", names[winner])
		}
	}
	fmt.Println("Goodbye.")
}

func getChoice(player string) int {
	for {
		var buf string

		fmt.Printf("%v, enter your choice\n", player)
		fmt.Println("Rock, Paper, Scissors, Lizard, Spock")
		fmt.Scanln(&buf)

		// Verify
		for k, v := range names {
			if v == buf {
				return k
			}
		}

		fmt.Println("You've made a invalid choice. Please try again")
	}
}
