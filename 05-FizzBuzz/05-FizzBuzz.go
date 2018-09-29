// 05: FizzBuzz
package main

import (
	"fmt"
)

// Print all the numbers from 1 to 100
// If it's divisible by 3 print Fizz, 5 print Buzz
// 15 print FizzBuzz otherwise print just the number
//
// How some people mess this up is beyond me
func main() {
	fmt.Println("05: Fizz Buzz")
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
