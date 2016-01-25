package hamming

import (
	"errors"
)

const TestVersion = 2

func Distance(a, b string) (hammingDistance int, err error) {
	hammingDistance = 0

	if len(a) != len(b) {
		return 0, errors.New("sequences of unequal length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
