package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

//RemoveSpaces removes spaces in number string
func RemoveSpaces(input string) string {
	noSpaces := strings.Replace(input, " ", "", -1)
	return noSpaces
}

//Reverse returns a reversed string
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

//DoubledDigits doubles the even numbered digits starting from the right of the string and returns the string with digits doubled while skiping non-digit characters
func DoubledDigits(input string) string {
	Doubled := ""
	Counter := 0
	for _, r := range reverse(input) {
		//if the rune is a digit, check its position
		if unicode.IsDigit(r) {
			//only iterate counter for digits
			Counter++
			//double every other digit from the right
			if Counter%2 == 0 {
				// converts rune to int and doubles
				digit := int(r-'0') * 2
				//test if doubled digit is greater than 9 and subtracts 9 if so
				if digit > 9 {
					digit = digit - 9
					Doubled += strconv.Itoa(digit)
				} else {
					Doubled += strconv.Itoa(digit)
				}
			} else {
				//add odd digits directly to Doubled
				Doubled += string(r)
			}
		} else {
			//returns an invalid string which triggers a false result in Valid
			Doubled = "1"
		}
	}
	return Doubled
}

//SumOfDigits returns an integer sum of all digits provided in input
func Sum(input string) int {
	sum := 0
	for _, r := range input {
		//converts digit rune to int and adds to sum
		digit := int(r - '0')
		sum += digit
	}
	return sum
}

//Valid determines whether a credit card number or CSIN is a valid number. Returns a boolean.
func Valid(input string) bool {
	var isValid bool
	if len(RemoveSpaces(input)) <= 1 {
		isValid = false
	} else {
		doubled := DoubledDigits(RemoveSpaces(input))
		sum := Sum(doubled)
		if sum%10 == 0 {
			isValid = true
		}
	}
	return isValid
}
