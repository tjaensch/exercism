package wordcount

import (
	"regexp"
	"strings"
)

const TestVersion = 1

type Frequency map[string]int

func normalize(text string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	normalized := reg.ReplaceAllString(text, " ")
	normalized = strings.ToLower(normalized)
	return normalized
}

func WordCount(phrase string) Frequency {
	phrase = normalize(phrase)
	result := make(Frequency)
	words := strings.Fields(phrase)
	for _, word := range words {
		_, ok := result[word]
		if ok == true {
			result[word] += 1
		} else {
			result[word] = 1
		}
	}
	return result
}
