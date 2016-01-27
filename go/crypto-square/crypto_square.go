package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const TestVersion = 1

func normalize(text string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	normalized := reg.ReplaceAllString(text, "")
	normalized = strings.ToLower(normalized)
	return normalized
}

func Encode(text string) string {
	text = normalize(text)
	numCols := int(math.Ceil(math.Sqrt(float64(len(text)))))
	cols := make([]string, numCols)
	for i, row := range text {
		cols[i%numCols] += string(row)
	}
	return strings.Join(cols, " ")
}
