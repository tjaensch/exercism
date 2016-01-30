package lsproduct

import (
	"errors"
	"strconv"
)

const TestVersion = 2

func LargestSeriesProduct(digits string, n int) (int, error) {
	if len(digits) < n {
		return 0, errors.New("digits < n")
	} else if n < 0 {
		return 0, errors.New("n cannot be negative")
	}

	var maxproduct int
	for i := 0; i <= len(digits)-n; i++ {
		substring := 1
		for ii := 0; ii < n; ii++ {
			n, err := strconv.Atoi(string(digits[ii+i]))
			if err != nil {
				return 0, errors.New("invalid digits")
			}
			substring *= n
		}
		if maxproduct <= substring {
			maxproduct = substring
		}
	}
	return maxproduct, nil
}
