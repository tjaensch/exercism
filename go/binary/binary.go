package binary

import (
	"errors"
)

func ParseBinary(bin string) (int, error) {
	val := 0
	for _, digit := range bin {
		switch digit {
		case '1':
			val = (val << 1) + 1
		case '0':
			val = (val << 1)
		default:
			return 0, errors.New("digit not 0 or 1")
		}
	}
	return val, nil
}
