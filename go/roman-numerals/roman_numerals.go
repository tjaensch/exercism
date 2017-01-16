package romannumerals

import (
	"errors"
)

const TestVersion = 1

func ToRomanNumeral(num int) (string, error) {

	if num <= 0 || num > 3999 {
		return "", errors.New("num out of range")
	}

	figure := []int{1000, 100, 10, 1}
	romanDigitA := []string{
		1:    "I",
		10:   "X",
		100:  "C",
		1000: "M",
	}
	romanDigitB := []string{
		1:    "V",
		10:   "L",
		100:  "D",
		1000: "MMMMM",
	}

	romanSlice := []string{}
	x := ""

	for _, f := range figure {
		digit, i, v := int(num/f), romanDigitA[f], romanDigitB[f]
		switch digit {
		case 1:
			romanSlice = append(romanSlice, string(i))
		case 2:
			romanSlice = append(romanSlice, string(i)+string(i))
		case 3:
			romanSlice = append(romanSlice, string(i)+string(i)+string(i))
		case 4:
			romanSlice = append(romanSlice, string(i)+string(v))
		case 5:
			romanSlice = append(romanSlice, string(v))
		case 6:
			romanSlice = append(romanSlice, string(v)+string(i))
		case 7:
			romanSlice = append(romanSlice, string(v)+string(i)+string(i))
		case 8:
			romanSlice = append(romanSlice, string(v)+string(i)+string(i)+string(i))
		case 9:
			romanSlice = append(romanSlice, string(i)+string(x))
		}

		num -= digit * f
		x = i
	}
	ret := ""
	for _, e := range romanSlice {
		ret += e
	}
	return ret, nil
}
