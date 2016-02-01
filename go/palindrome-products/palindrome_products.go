package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalindrome(num int) bool {
	numString := strconv.Itoa(num)
	for i, j := 0, len(numString)-1; i < j; i, j = i+1, j-1 {
		if numString[i] != numString[j] {
			return false
		}
	}
	return true
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return Product{}, Product{}, err
	}
	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			p := x * y
			if !isPalindrome(p) {
				continue
			}
			compare := func(i *Product, check bool) {
				switch {
				case i.Factorizations == nil || check:
					*i = Product{p, [][2]int{{x, y}}}
				case p == i.Product:
					i.Factorizations =
						append(i.Factorizations, [2]int{x, y})
				}
			}
			compare(&pmin, p < pmin.Product)
			compare(&pmax, p > pmax.Product)
		}
	}
	if len(pmax.Factorizations) == 0 {
		err = errors.New("No palindromes")
	}
	return
}
