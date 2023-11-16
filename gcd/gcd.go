package practice

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
