package main

import (
	"fmt"
	"os"
)

func Rostring(s string) {
	words := []string{}
	word := ""

	for _, char := range s {
		if char != ' ' {
			word += string(char)
		} else if char == ' ' {
			words = append(words, word)
			word = ""
			word += string(char)
			words = append(words, word)
			word = ""
		}
	}
	words = append(words, word)
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
	}
	fmt.Println()
}

func main() {
	args := os.Args[1:]
	Rostring(args[0])

	// "listen", "silent" silent
}
