//Package twofer implements a function that takes a varibale and places it in a sentence.
package twofer

// ShareWith returns string "one for (name or you), one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
