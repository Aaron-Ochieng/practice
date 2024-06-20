package main

func Atoibase2(s, b string) int {
	baselen := len(b)
	index := 0
	res := 0
	for _, char := range s {
		for pos, base := range b {
			if char == base {
				index = pos
			}
		}
		res = res*baselen + index
	}
	return res
}
