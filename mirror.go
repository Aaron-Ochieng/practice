package practice

func Mirror(s string) string {
	arr := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' || c <= 'z' {
			arr[i] = 219 - c
		} else if c >= 'A' || c <= 'Z' {
			arr[i] = 155 - c
		} else {
			arr[i] = c
		}
	}
	return string(arr)
}

// "a"
