package pythagorean

type Triplet [3]int

func Range(min, max int) (triplet []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					triplet = append(triplet, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

func Sum(sum int) (triplet []Triplet) {
	max := sum / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := sum - a - b; a*a+b*b == c*c {
				triplet = append(triplet, Triplet{a, b, c})
			}
		}
	}
	return
}
