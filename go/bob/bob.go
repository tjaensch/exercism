package bob

import "strings"

const TestVersion = 1

func Hey(what string) string {
	what = strings.TrimSpace(what)
	if what == "" {
		return "Fine. Be that way!"
	}
	if what == strings.ToUpper(what) && what != strings.ToLower(what) {
		return "Whoa, chill out!"
	}
	if (what[len(what)-1]) == '?' {
		return "Sure."
	}
	return "Whatever."
}
