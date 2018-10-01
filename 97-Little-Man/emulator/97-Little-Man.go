// 97: Little Man Computer
package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var menMap = map[int]string{
	100: "ADD",
	200: "SUB",
	300: "STA",
	400: "LDA",
	500: "BRA",
	600: "BRZ",
	700: "BRP",
	800: "INP",
	901: "OUT",
	902: "HLT",
	000: "DAT",
}

type littleMan struct {
	PC          int
	Memory      [100]int
	Inbox       int
	Outbox      int
	Accumulator int
}

var stdin *bufio.Reader
var stdout *bufio.Writer

var DEBUG bool
var VERBOSE bool

func init() {
	var showHelp bool
	// Set up buffered IO
	stdin = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)

	// Parse flags
	flag.BoolVar(&DEBUG, "d", false, "Dumps state on Halt")
	flag.BoolVar(&VERBOSE, "v", false, "Dumps state after every step")
	flag.BoolVar(&showHelp, "h", false, "Show help")
	flag.Parse()

	// Display usage and exit
	if len(flag.Args()) != 1 || showHelp {
		fmt.Println("Usage: ./97 program.hex")
		fmt.Println("Input is read from STDIN, Output written on exit")
		flag.PrintDefaults()
		return
	}
}

func main() {
	fmt.Println("03: Little Man Computer Simulator")
	var comp littleMan

	// Load pgm into computer
	comp.LoadProgram(flag.Arg(0))
	// Exec loop
	for {
		isHalted, err := comp.Step()
		if VERBOSE {
			comp.Dump()
		}

		if isHalted {
			break
		}

		if err != nil {
			panic(err)
		}

		// Give a bit to watch the display
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Execution Terminated.")
	if DEBUG || VERBOSE {
		comp.Dump()
	}
}

// Load bytecode to memory
func (l *littleMan) LoadProgram(path string) (e error) {
	// Get raw data
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read file: %v", err)
		return err
	}

	// Decode to parsable format
	var pgm []int
	buf := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&pgm)
	if err != nil {
		fmt.Printf("Error decoding file: %v", err)
		return err
	}

	// Trim and store
	for i, v := range pgm {
		if i == 100 {
			fmt.Println("Warning! Program is larger then avalible memory")
			break
		}

		l.Memory[i] = v
	}

	return
}

// Advance execution
func (l *littleMan) Step() (isHalt bool, err error) {
	// Retrieve the current ins
	ins, data := parseCode(l.Memory[l.PC])

	// Advance the PC
	if l.PC == 100 {
		return false, errors.New("pc pointed out of memory. Halting")
	}
	l.PC++

	// Handle instruction
	switch ins {
	case "ADD":
		l.Accumulator += l.Memory[data]
	case "SUB":
		l.Accumulator -= l.Memory[data]
	case "STA":
		l.Memory[data] = l.Accumulator
	case "LDA":
		l.Accumulator = l.Memory[data]
	case "BRA":
		l.PC = data
	case "BRZ":
		if l.Accumulator == 0 {
			l.PC = data
		}
	case "BRP":
		if l.Accumulator != 0 {
			l.PC = data
		}
	case "INP":
		fmt.Printf("Prompting for input:")
		c, err := stdin.ReadByte()
		if err != nil {
			fmt.Printf("Error reading input. %v", err)
			fmt.Println("Halting")
			return true, err
		}

		l.Accumulator = int(c)
	case "OUT":
		err = stdout.WriteByte(byte(l.Accumulator))
		if err != nil {
			fmt.Printf("Error writig output. %v", err)
			fmt.Println("Halting")
			isHalt = true
		}
	case "HLT":
		isHalt = true
	}
	return
}

func (l *littleMan) Dump() {
	fmt.Println("System Dump: ")
	fmt.Printf("PC: %v Inbox: %v Outbox: %v Acc: %v\n\n", l.PC, l.Inbox, l.Outbox, l.Accumulator)
	fmt.Println("Memory Dump:")

	// Pretty hex dump the core mem
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ] ", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("%04x ", l.Memory[i*10+j])
		}
		fmt.Printf("\n")
	}
}

// Seperate the instruction from its operand
func parseCode(code int) (ins string, data int) {
	data = code % 100
	switch code - (code % 100) {
	case 100:
		ins = "ADD"
	case 200:
		ins = "SUB"
	case 300:
		ins = "STA"
	case 400:
		ins = "LDA"
	case 500:
		ins = "BRA"
	case 600:
		ins = "BRZ"
	case 700:
		ins = "BRP"
	case 800:
		ins = "INP"
	case 901:
		ins = "OUT"
	case 902:
		ins = "HLT"
	case 000:
		ins = "DAT"
	}
	return
}
