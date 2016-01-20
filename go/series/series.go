package slice

func All(n int, s string) []string {

	substring := make([]string, len(s)-n+1)

	if n > len(s) {
		return nil
	}

	for i := 0; i <= len(s)-n; i++ {
		substring[i] = s[i : i+n]
	}
	return substring

}

func Frist(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
