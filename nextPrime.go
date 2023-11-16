package practice

func NextPrime(n int) int {
	for !IsPrime(n) {
		n++
	}
	return n
}
