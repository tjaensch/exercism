package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	m := FreqMap{}

	ch := make(chan FreqMap)
	for _, text := range texts {
		go func(text string) { ch <- Frequency(text) }(text)
	}
	for _ = range texts {
		for i, r := range <-ch {
			m[i] += r
		}
	}
	return m
}
