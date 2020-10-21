package luhn

import (
	"strconv"
)

//RemoveSpaces removes spaces in number string
func RemoveSpaces(input string) string {
	noSpaces := ""
	for _, r := range input {
		if r == ' ' {
			noSpaces += ""
		} else {
			noSpaces += string(r)
		}
	}
	return noSpaces
}

//DoubledDigits doubles the even numbered digits starting from the right of the string and returns the string with digits doubled
func DoubledDigits(input string) string {
	Doubled := ""
	for i, r := range input {
		if i-1%2 == 0 {
			digit := int(r-'0') * 2
			if digit >= 9 {
				digit = digit - 9
				Doubled += strconv.Itoa(digit)
			} else {
				Doubled += strconv.Itoa(digit)
			}
		} else {
			Doubled += string(r)
		}
	}
	return Doubled
}

//SumOfDigits returns an integer sum of all digits provided in input
func SumOfDigits(input string) int {
	sum := 0
	for _, r := range input {
		digit := int(r - '0')
		sum += digit
	}
	return sum
}

//Valid determines whether a credit card number or CSIN is a valid number. Returns a boolean.
func Valid(input string) bool {
	var isValid bool
	if len(input) <= 1 {
		isValid = false
	} else {
		doubled := DoubledDigits(input)
		sum := SumOfDigits(doubled)
		if sum%10 == 0 {
			isValid = true
		}
	}
	return isValid
}
