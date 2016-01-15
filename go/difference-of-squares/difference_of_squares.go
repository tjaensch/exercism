package diffsquares

func SquareOfSums(n int) int {
	var valueSquareOfSums int
	for i := 1; i <= n; i++ {
		valueSquareOfSums += i
	}
	return valueSquareOfSums * valueSquareOfSums
}

func SumOfSquares(n int) int {
	var valueSumOfSquares int
	for i := 1; i <= n; i++ {
		valueSumOfSquares += i * i
	}
	return valueSumOfSquares
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
