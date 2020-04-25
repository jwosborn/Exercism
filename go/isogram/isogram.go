package isogram

import "strings"

//IsIsogram takes a string and returns true if the word is an isogram and false if not.
func IsIsogram(a string) bool {
	n := true
	//remove spaces
	word := strings.Replace(a, " ", "", -1)
	//remove hyphens
	if strings.Contains(a, "-") {
		word = strings.Replace(a, "-", "", -1)
	}
	//count letter occurances
	for _, r := range word {
		if strings.Count(strings.ToLower(word), strings.ToLower(string(r))) >= 2 {
			n = false
		}
	}
	return n
}
