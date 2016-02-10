package strand

import (
	"strings"
)

const TestVersion = 2

func ToRNA(dna string) string {

	dnaSlice := strings.Split(dna, "")
	rnaSlice := make([]string, 0)

	for _, v := range dnaSlice {
		switch v {
		case "G":
			v = "C"
			rnaSlice = append(rnaSlice, v)
		case "C":
			v = "G"
			rnaSlice = append(rnaSlice, v)
		case "T":
			v = "A"
			rnaSlice = append(rnaSlice, v)
		case "A":
			v = "U"
			rnaSlice = append(rnaSlice, v)
		}
	}
	return strings.Join(rnaSlice, "")
}
