package triangle

import "math"

// Kind is a string specific to triangles
type Kind string

const (
	//NaT not a triangle
	NaT Kind = "NaT"
	//Equ equilateral
	Equ Kind = "Equ"
	//Iso isosceles
	Iso Kind = "Iso"
	//Sca scalene
	Sca Kind = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if (a == b && a != c) || (a != b && a == c) && (b != c) {
		k = Iso
	} else if a != b && a != c && b == c {
		k = Iso
	} else if a == b && a == c {
		k = Equ
	} else if a != b && b != c && a != c {
		k = Sca
	}

	return k
}
