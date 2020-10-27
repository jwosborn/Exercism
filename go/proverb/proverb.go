// Package proverb returns a proverb given a list of strings
package proverb

// Proverb returns a proverb given a list of strings
func Proverb(rhyme []string) []string {
	var proverb []string
	switch len(rhyme) {
	case 0:
		return proverb
	case 1:
		proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
	case 2:
		proverb = append(proverb, ("For want of a " + rhyme[0] + " the " + rhyme[1] + " was lost."), ("And all for the want of a " + rhyme[0] + "."))
	default:
		for i := range rhyme {
			if i == len(rhyme)-2 {
				proverb = append(proverb, ("For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."), ("And all for the want of a " + rhyme[0] + "."))
			} else if i == len(rhyme)-1 {
				return proverb
			} else {
				proverb = append(proverb, ("For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."))
			}
		}
	}
	return proverb
}
