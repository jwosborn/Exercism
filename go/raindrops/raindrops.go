//Package raindrops provides a function Convert that takes an integer and returns a string
package raindrops

import "strconv"

//Convert takes an integer and returns a string
func Convert(a int) string {
	var message string
	if (a % 3) == 0 {
		message = message + "Pling"
	} else {
		message = message
	}
	if (a % 5) == 0 {
		message = message + "Plang"
	} else {
		message = message
	}
	if (a % 7) == 0 {
		message = message + "Plong"
	} else {
		message = message
	}
	if message == "" {
		return strconv.Itoa(a)
	}
	return message
}
