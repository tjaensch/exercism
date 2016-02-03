package brackets

const TestVersion = 2

type bracket int

const (
	open bracket = iota
	closed
	other
)

var brackets = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func checkBracket(char rune) bracket {
	for k, v := range brackets {
		switch char {
		case k:
			return open
		case v:
			return closed
		}
	}
	return other
}

func Bracket(input string) (bool, error) {
	var matches []rune
	for _, v := range input {
		switch checkBracket(v) {
		case open:
			matches = append(matches, brackets[v])
		case closed:
			if 0 < len(matches) && matches[len(matches)-1] == v {
				matches = matches[:len(matches)-1]
			} else {
				return false, nil
			}
		}
	}
	return len(matches) == 0, nil
}
