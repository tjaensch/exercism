package queenattack

import (
	"errors"
)

//check if string combination is a valid chess square
func validSquare(f string) error {
	if len(f) != 2 || f[0] < 'a' || f[0] > 'h' || f[1] > '8' {
		return errors.New("Invalid square")
	}
	return nil
}

func CanQueenAttack(w, b string) (attack bool, err error) {
	if err = validSquare(w); err != nil {
		return
	}
	if err = validSquare(b); err != nil {
		return
	}

	if w[0]-b[0] == 0 && w[1]-b[1] == 0 {
		err = errors.New("Same square")
	} else if w[0]-b[0] == 0 || w[1]-b[1] == 0 || w[0]-b[0] == w[1]-b[1] || w[0]-b[0] == -(w[1]-b[1]) {
		attack = true
	}
	return
}
