package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid square")
	}
	// left shift equals multiplying by 2
	return 1 << (uint64(n) - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
