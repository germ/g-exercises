// 02: Temperture Converter
package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		// Use Kelvin for internal storage
		// This is a long function, yet readable
		// I think the word is quaint
		var tempK, fromTemp, toTemp float64
		var buf, fromUnit, toUnit string

		fmt.Println("02: Temperture Converter")
		for {
			fmt.Println("What unit is the temp in?\n(K)elvin, (C)elsius or (F)ahrenheit")
			fmt.Scanln(&buf)

			// Trying something new here, not sure if it will stick
			switch buf {
			case "K":
			case "F":
			case "C":
			default:
				fmt.Println("Invalid input, Try that again bud.")
				continue
			}

			fromUnit = buf
			break
		}

		// If being reused, break this out
		for {
			fmt.Println("What unit Are we converting to?\n(K)elvin, (C)elsius or (F)ahrenheit")
			fmt.Scanln(&buf)

			// Trying something new here, not sure if it will stick
			switch buf {
			case "K":
			case "F":
			case "C":
			default:
				fmt.Println("Invalid input, Try that again bud.")
				continue
			}

			toUnit = buf
			break
		}

		// Some code actually does not need comments everywhere
		// As long as you're not being clever anyway
		for {
			fmt.Println("Lastly, What's the temperture?")
			fmt.Scanln(&buf)

			var err error
			fromTemp, err = strconv.ParseFloat(buf, 64)
			if err != nil {
				fmt.Println("Invalid input there bud.")
				continue
			}

			break
		}

		// Get that conversion on
		// Convert to Kelvin
		switch fromUnit {
		case "K":
			tempK = fromTemp
		case "C":
			tempK = fromTemp + 273
		case "F":
			tempK = (fromTemp + 459.67) * (5.0 / 9.0)
		}

		// Convert to final unit
		switch toUnit {
		case "K":
			toTemp = tempK
		case "F":
			toTemp = tempK*(9.0/5.0) - 459.67
		case "C":
			toTemp = tempK - 273
		}

		// Display and prompt
		fmt.Printf("%0.2f%v is %0.2f%v\n", fromTemp, fromUnit, toTemp, toUnit)

		// Get out if we can
		fmt.Println("Would you like to do another conversion? (Y/N)")
		fmt.Scanln(&buf)
		if buf == "N" {
			fmt.Println("Goodbye.")
			return
		}
	}
}
