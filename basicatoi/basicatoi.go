package practice

func BasicAtoi(s string) int {
	result := 0
	//
	for _, char := range s {
		if char < '0' && char > '9' {
			return 0
		} else {
			result = (result * 10) + int(char-'0')
		}
	}
	return result
}
