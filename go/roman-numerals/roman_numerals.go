//package romannumerals contains a function that takes an arabic integer and returns a roman numeral string of equivalent value
package romannumerals

import (
	"fmt"
)

func ToRomanNumeral(n int) (string, error) {
	roman := ""
	thousands := n / 1000
	hundreds := n % 1000 / 100
	tens := n % 1000 % 100 / 10
	ones := ((n % 1000) % 100) % 10
	if n < 0 && n >= 3001 {
		return "", fmt.Errorf("bad bad")
	} else {
		for n != 0 {
			for thousands < 4 && thousands > 0 {
				roman += "M"
				thousands--
				n -= 1000
			}
			for hundreds == 9 {
				roman += "CM"
				hundreds -= 9
				n -= 900
			}
			for hundreds == 5 {
				roman += "D"
				hundreds -= 5
				n -= 500
			}
			for hundreds == 4 {
				roman += "CD"
				hundreds -= 4
				n -= 400
			}
			for hundreds > 5 || hundreds < 4 && hundreds > 0 {
				roman += "C"
				hundreds--
				n -= 100
			}
			for tens == 9 {
				roman += "XC"
				tens -= 9
				n -= 90
			}
			for tens > 5 || tens < 4 && tens > 0 {
				roman += "X"
				tens--
				n -= 10
			}
			for tens == 5 {
				roman += "L"
				tens -= 5
				n -= 50
			}
			for tens == 4 {
				roman += "IL"
				tens -= 4
				n -= 40
			}
			for ones == 9 {
				roman += "IX"
				ones -= 9
				n -= 9
			}
			for ones > 5 || ones < 4 && ones > 0 {
				roman += "I"
				ones--
				n--
			}
			for ones == 5 {
				roman += "V"
				ones -= 5
				n -= 5
			}
			for ones == 4 {
				roman += "IV"
				ones -= 4
				n -= 4
			}

		}
	}
	return roman, nil
}
