package main

func main() {
	s := "helloiEAthura"
	result := rotateVowels(s)
	println(s)
	println(result)
}

func rotateVowels(s string) string {
	sb := []byte(s)

	mid := len(sb) / 2
	if len(s)%2 == 1 {
		mid++
	}
	index := len(sb) - 1

	for i := 0; i <= mid; i++ {
		if isVowel(sb[i]) {
			for j := index; j > i; j-- {
				if isVowel(sb[j]) {
					sb[i], sb[j] = sb[j], sb[i]
					index = j - 1
					break
				}
				index--
			}
		}
	}
	return string(sb)
}

func isVowel(ch byte) (isvowel bool) {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		isvowel = true
	}
	return
}
