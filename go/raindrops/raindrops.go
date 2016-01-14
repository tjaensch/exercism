package raindrops

import "strconv"

const TestVersion = 1

func Convert(num int) string {
	if num%3 == 0 && num%5 == 0 && num%7 == 0 {
		return "PlingPlangPlong"
	} else if num%3 == 0 && num%5 == 0 {
		return "PlingPlang"
	} else if num%3 == 0 && num%7 == 0 {
		return "PlingPlong"
	} else if num%5 == 0 && num%7 == 0 {
		return "PlangPlong"
	} else if num%3 == 0 {
		return "Pling"
	} else if num%5 == 0 {
		return "Plang"
	} else if num%7 == 0 {
		return "Plong"
	} else {
		return strconv.Itoa(num)
	}

}
