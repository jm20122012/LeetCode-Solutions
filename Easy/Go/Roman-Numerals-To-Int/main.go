// Given a roman numeral, convert it to an integer.

package main

import (
	"fmt"
	"strings"
)

var numerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var sum int = 0
	var index int = 0
	strSlice := strings.Split(s, "")

	for index = 0; index < len(strSlice); index++ {
		var val string = strSlice[index]

		var nextVal string = ""
		if index != len(strSlice)-1 {
			nextVal = strSlice[index+1]
		}
		if val+nextVal == "IV" {
			fmt.Printf("[Special Case] Adding 4 to %d - New sum: %d\n", sum, sum+4)
			sum += 4
			index += 1
		} else if val+nextVal == "IX" {
			fmt.Printf("[Special Case] Adding 9 to %d - New sum: %d\n", sum, sum+9)
			sum += 9
			index += 1
		} else if val+nextVal == "XL" {
			fmt.Printf("[Special Case] Adding 40 to %d - New sum: %d\n", sum, sum+40)
			sum += 40
			index += 1
		} else if val+nextVal == "XC" {
			fmt.Printf("[Special Case] Adding 90 to %d - New sum: %d\n", sum, sum+90)
			sum += 90
			index += 1
		} else if val+nextVal == "CD" {
			fmt.Printf("[Special Case] Adding 400 to %d - New sum: %d\n", sum, sum+400)
			sum += 400
			index += 1
		} else if val+nextVal == "CM" {
			fmt.Printf("[Special Case] Adding 900 to %d - New sum: %d\n", sum, sum+900)
			sum += 900
			index += 1
		} else {
			fmt.Printf("Adding %s to %d - New sum: %d\n", val, sum, sum+numerals[val])
			sum += numerals[val]
		}
	}
	return sum
}

func main() {
	fmt.Print(romanToInt("MCMXCIV"))
}
