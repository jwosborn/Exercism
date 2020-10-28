//strand contains a function that takes a dna string and returns an rna string
package strand

//ToRNA takes a dna string and returns an RNA string
func ToRNA(dna string) string {
	rna := ""
	for _, r := range dna {
		switch r {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
