//Package raindrops provides a function Convert that takes an integer and returns a string
package raindrops

import "strconv"

//Convert takes an integer and returns a string
func Convert(a int) string {
	var message string
	if (a % 3) == 0 {
		message += "Pling"
	}
	if (a % 5) == 0 {
		message += "Plang"
	}
	if (a % 7) == 0 {
		message += "Plong"
	}
	if message == "" {
		return strconv.Itoa(a)
	}
	return message
}
