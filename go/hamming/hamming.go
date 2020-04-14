//Package hamming provides a function Distance that gives the amount of differences between two DNA strings
package hamming

import (
	"errors"
)

//Distance calculates the amount of differences between two strings and returns an int
func Distance(a, b string) (int, error) {
	//Make sure strings are the same length
	if len(a) != len(b) {
		return 0, errors.New("the lengths of these two strings do not match")
	}
	//start distance at 0
	d := 0
	//test for differences in strings based on index
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
