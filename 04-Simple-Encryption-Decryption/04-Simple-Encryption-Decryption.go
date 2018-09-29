// 04: Simple Encryption-Decryption
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Rather then reimplementing RSA or something, this program
// Uses a One Time Pad of arbitrary data (In this case a random
// Image) that is passed. Typically this would be a photo that is
// byte identical, but could be any shared file. If the message is
// longer then the file supplied, it is repeated.
// The message is dumped into filename.enc after
//
// Because this is symetric encryption and decryption are the same
func main() {
	fmt.Println("04: Simple Encryption-Decryption")
	var key, plainText, cipherText []byte
	var err error

	// Print help
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./04 padFile messageFile")
		return
	}

	// Try reading files into memory
	key, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Could not pad: ", err)
		return
	}
	plainText, err = ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("Could not message: ", err)
		return
	}

	// Loop and XOR!
	cipherText = make([]byte, len(plainText))
	for i := 0; i < len(cipherText); i++ {
		cipherText[i] = plainText[i] ^ key[i%len(key)]
	}

	// Check for existing suffix and strip
	var outName string
	if strings.HasSuffix(os.Args[2], ".enc") {
		outName = os.Args[2][:len(os.Args[2])-4]
	} else {
		outName = os.Args[2] + ".enc"
	}

	// Write it!
	err = ioutil.WriteFile(outName, cipherText, 0755)
	if err != nil {
		fmt.Println("Could not write file: ", err)
	} else {
		fmt.Println("File sucessfully written!", outName)
	}
}
