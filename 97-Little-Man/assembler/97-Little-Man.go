// 97: Little Man Computer
package main

import (
	"fmt"
	"os"
)

type LittleMan struct {
	PC          byte
	Memory      [100]byte
	Inbox       byte
	Outbox      byte
	Accumulator byte
}

func main() {
	fmt.Println("03: Little Man Computer Simulator")
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./97 program.hex")
		return
	}

}

// Load bytecode to memory
func (l *LittleMan) LoadProgram(pgm []byte) {
}

// Advance execution
func (l *LittleMan) Step() {
}

// Place byte into Inbox
func (l *LittleMan) Input() {
}

// Retrieve byte from Inbox
func (l *LittleMan) Output() {
}
