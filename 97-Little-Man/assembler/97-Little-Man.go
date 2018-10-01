// 97: Little Man Computer
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

var menMap = map[string]int{
	"ADD": 100,
	"SUB": 200,
	"STA": 300,
	"LDA": 400,
	"BRA": 500,
	"BRZ": 600,
	"BRP": 700,
	"INP": 800,
	"OUT": 901,
	"HLT": 902,
	"DAT": 000,
}

func main() {
	fmt.Println("03: Little Man Computer Assembler")
	var byteCode []int

	// Did we get run properly?
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./97 program.lmc")
		fmt.Println("Creates fileName.hex")
		return
	}

	// Read entire file and split
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	byteCode = assembleSource(string(f))
	// Convert to easily parseable format (gob)
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(byteCode)

	// Write our output
	filePath := path.Base(os.Args[1]) + ".hex"
	err = ioutil.WriteFile(filePath, buf.Bytes(), 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Compilation sucessful, byteCode written to %v", filePath)
}

// Read the source and return bytecode suitable for running
func assembleSource(srcIn string) (byteCode []int) {
	// So there's a few things we need to do. The first is to strip any comments or blank
	// lines out. Then we scan over the source and make a list of the address of any labels
	// we then substitute any labels with those offsets, lastly we convert the generated code
	// Into bytecode. Woo!

	// Step one: Remove comments
	src := strings.Split(srcIn, "\n")
	label := make(map[string]int)

	var tmp []string
	for _, v := range src {
		// Blank line
		if len(v) == 0 {
			continue
		}

		// Comment
		if strings.HasPrefix(v, ";") {
			continue
		}

		tmp = append(tmp, v)
	}
	src = tmp

	// Step 2: Locate any labels, these will be in position one and not a instruction
	for i, v := range src {
		f := strings.Fields(v)
		// Check against memnomics
		if _, ok := menMap[f[0]]; !ok {
			label[f[0]] = i
		}
	}

	// Step 3: Loop over removing label defs and replacing refrences
	var processedSrc [][]string
	for _, v := range src {
		line := strings.Fields(v)

		// Remove the label def if found
		if _, ok := label[line[0]]; ok {
			line = line[1:]
		}

		//check operand for label and sub
		if len(line) > 1 {
			if v, ok := label[line[1]]; ok {
				line[1] = fmt.Sprintf("%d", v)
			}
		}

		processedSrc = append(processedSrc, line)
	}

	// Step 4: Convert to bytecode, 1 byte per ins/datum
	for i, v := range processedSrc {
		// Retrieve instruction code
		ins, ok := menMap[v[0]]
		if !ok {
			fmt.Printf("Error Processing line %v: %v", i, v)
			panic("")
		}

		// Handle dataless instructions
		if v[0] == "INP" || v[0] == "OUT" || v[0] == "HLT" {
			byteCode = append(byteCode, ins)
			continue
		}

		// Trim data and add
		data, err := strconv.Atoi(v[1])
		if err != nil {
			fmt.Printf("Invalid data on line %v: %v ", i, err)
			panic("")
		}
		data = ins + data

		byteCode = append(byteCode, data)
	}
	return
}
