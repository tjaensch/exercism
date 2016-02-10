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
	roman_digitA := []string{
		1:    "I",
		10:   "X",
		100:  "C",
		1000: "M",
	}
	roman_digitB := []string{
		1:    "V",
		10:   "L",
		100:  "D",
		1000: "MMMMM",
	}

	roman_slice := []string{}
	x := ""

	for _, f := range figure {
		digit, i, v := int(num/f), roman_digitA[f], roman_digitB[f]
		switch digit {
		case 1:
			roman_slice = append(roman_slice, string(i))
		case 2:
			roman_slice = append(roman_slice, string(i)+string(i))
		case 3:
			roman_slice = append(roman_slice, string(i)+string(i)+string(i))
		case 4:
			roman_slice = append(roman_slice, string(i)+string(v))
		case 5:
			roman_slice = append(roman_slice, string(v))
		case 6:
			roman_slice = append(roman_slice, string(v)+string(i))
		case 7:
			roman_slice = append(roman_slice, string(v)+string(i)+string(i))
		case 8:
			roman_slice = append(roman_slice, string(v)+string(i)+string(i)+string(i))
		case 9:
			roman_slice = append(roman_slice, string(i)+string(x))
		}

		num -= digit * f
		x = i
	}
	ret := ""
	for _, e := range roman_slice {
		ret += e
	}
	return ret, nil
}
