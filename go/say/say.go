package say

var small = []string{"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
	"fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninety"}
var scale = []string{"thousand", "million", "billion",
	"trillion", "quadrillion", "quintillion"}

func Say(num uint64) string {
	switch {
	case num < 20:
		return small[num]
	case num < 100:
		smallNum := tens[num/10]
		s := num % 10
		if s > 0 {
			smallNum += "-" + small[s]
		}
		return smallNum
	case num < 1000:
		midNum := small[num/100] + " hundred"
		s := num % 100
		if s > 0 {
			midNum += " " + Say(s)
		}
		return midNum
	}
	scaleNum := ""
	if p := num % 1000; p > 0 {
		scaleNum = Say(p)
	}
	for i := 0; num >= 1000; i++ {
		num /= 1000
		if p := num % 1000; p > 0 {
			endNum := Say(p) + " " + scale[i]
			if scaleNum > "" {
				endNum += " " + scaleNum
			}
			scaleNum = endNum
		}
	}
	return scaleNum
}
