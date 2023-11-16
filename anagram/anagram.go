package main

import "fmt"

func sort(s string) string {
	srune := []rune(s)
	for i := 0; i < len(srune); i++ {
		for j := i + 1; j < len(srune); j++ {
			if srune[i] > srune[j] {
				srune[i], srune[j] = srune[j], srune[i]
			}
		}
	}
	return string(srune)
}

func StringToRune(s string) string {
	word := ""
	for _, char := range s {
		if char != ' ' {
			word += string(char)
		} else if char == ' ' {
			continue
		}
	}
	return word
}

func Anagram(str, comp string) bool {
	a := StringToRune(str)
	b := StringToRune(comp)

	sA := sort(a)
	sB := sort(b)

	return sA == sB
}

func main() {
	fmt.Println(Anagram("listen", "silent"))
	fmt.Println(Anagram("alem", "school"))
	fmt.Println(Anagram("neat", "a net"))
	fmt.Println(Anagram("anna madrigal", "a man and a girl"))
}
