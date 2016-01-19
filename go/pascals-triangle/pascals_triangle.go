package pascal

func NextRow(prevRow []int) []int {
	nextRow := make([]int, len(prevRow)+1, len(prevRow)+1)
	nextRow[0], nextRow[len(prevRow)] = 1, 1
	for i := 0; i < len(prevRow)-1; i++ {
		nextRow[i+1] = prevRow[i] + prevRow[i+1]
	}
	return nextRow
}

func Triangle(n int) [][]int {
	row := [][]int{}
	for i := 0; i < n; i++ {
		if i == 0 {
			row = append(row, []int{1})
		} else {
			row = append(row, NextRow(row[i-1]))
		}
	}
	return row
}
