package luhn

func get(num string) []int {
	numSlice := make([]int, 0, len(num))
	for _, el := range num {
		if el >= '0' && el <= '9' {
			numSlice = append(numSlice, int(el-'0'))
		}
	}
	return numSlice
}

func checkSum(numSlice []int) int {
	for i := len(numSlice) - 1; i >= 0; i -= 2 {
		el := 2 * numSlice[i]
		if el > 9 {
			el -= 9
		}
		numSlice[i] = el
	}
	sum := 0
	for _, el := range numSlice {
		sum += el
	}
	return sum
}

func Valid(num string) bool {
	numSlice := get(num)
	if len(numSlice) == 0 {
		return false
	}
	last := len(numSlice) - 1
	return (checkSum(numSlice[:last])+numSlice[last])%10 == 0
}

func AddCheck(num string) string {
	return num + string('0'+(10-checkSum(get(num))%10)%10)
}
