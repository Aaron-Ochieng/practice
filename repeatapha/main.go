package main

import "os"

func repeatAlpha(s string) {
	res := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			for i := 0; i <= int(char-'a'); i++ {
				res += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			for i := 0; i <= int(char-'A'); i++ {
				res += string(char)
			}
		} else {
			res += string(char)
		}
	}

	os.Stdout.WriteString(res + "\n")
}

func main() {
	args := os.Args[1:]
	repeatAlpha(args[0])
}
