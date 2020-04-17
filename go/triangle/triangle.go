package triangle

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
	count := 0
	sides := [3]float64{a, b, c}
	if a > 0 || b > 0 || c > 0 {
		k = NaT
	}
	for i := 0; i < len(sides); i++ {
		side := sides[i]
		if side == a || side == b || side == c {
			count++
		}
	}

	if count-1 == 1 {
		k = Sca
	} else if count-1 == 2 {
		k = Iso
	} else if count == 3 {
		k = Equ
	}

	return k
}
