package etl

import (
	"strings"
)

type oldFormat map[int][]string
type newFormat map[string]int

func Transform(input oldFormat) (output newFormat) {
	output = make(newFormat)
	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return
}
