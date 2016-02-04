package anagram

import (
	"sort"
	"strings"
)

func sortLetters(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func isAnagram(subject, candidate string) bool {
	return subject != candidate && sortLetters(subject) == sortLetters(candidate)
}

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var matches []string
	for _, c := range candidates {
		c = strings.ToLower(c)
		if isAnagram(subject, c) {
			matches = append(matches, c)
		}
	}
	return matches
}
