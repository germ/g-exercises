// 03: Age In Seconds (Including Leap)
package main

import (
	"fmt"
	"time"
)

func main() {
	var birthTime time.Time

	fmt.Println("03: Age In Seconds (Including Leap)")
	for {
		var buf string
		var err error

		fmt.Println("Please enter your birth date (DD/MM/YYYY)")
		fmt.Scanln(&buf)
		birthTime, err = time.Parse("02/01/2006", buf)

		// Break if we have a valid date
		if err == nil {
			break
		}
		fmt.Printf("That date is invalid: %v\n", err)
	}

	// Calc the diff and print
	diff := time.Now().Sub(birthTime)
	fmt.Printf("Since %v %.0f seconds have elapsed!\n", birthTime.Format("02.01.2006"), diff.Seconds())
	fmt.Println("Goodbye.")
}
