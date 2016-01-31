package sieve

func Sieve(limit int) (primes []int) {
	numbers := make([]bool, limit)
	for num := 2; num < limit; {
		for i := num + num; i < limit; i += num {
			numbers[i] = true
		}
		//avoid infinite loop
		for num++; num < limit && numbers[num]; num++ {
		}
	}
	for i := 2; i < limit; i++ {
		if !numbers[i] {
			primes = append(primes, i)
		}
	}
	return
}
