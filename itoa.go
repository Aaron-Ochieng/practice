package practice

func Itoa(n int) string {
	var isnegative bool
	str := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		isnegative = true
		n = -n
	}
	for n > 0 {
		digit := n % 10
		r := rune(digit + '0')
		str = string(r) + str
		n = n / 10
	}
	if isnegative {
		str = "-" + str
	}
	return str
}
