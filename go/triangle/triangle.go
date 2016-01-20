package triangle

import "math"

type Kind float64

const (
	Equ Kind = iota // equilateral
	Iso             // isosceles
	Sca             // scalene
	NaT             // not a triangle
)

func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	} else if a < 0 || b < 0 || c < 0 {
		return NaT
	} else if a+b <= c || b+c <= a || a+c <= b {
		return NaT
	} else if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
