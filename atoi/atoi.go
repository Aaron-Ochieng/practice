package practice

func Atoi(s string) int {
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[0] == '-' {
			sign = -1
			continue
		} else if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = (res * 10) + int(s[i]-'0')
	}
	return sign * res
}
