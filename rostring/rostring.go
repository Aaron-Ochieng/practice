package main

import (
	"fmt"
	"os"
)

func RoString(str string) {
	words := []string{}
	word := ""
	runes := []rune(str + " ")
	for i := 0; i < len(runes); i++ {
		if runes[i] != ' ' && i < len(runes)-1 {
			word += string(runes[i])
		} else if runes[i] == ' ' && word != "" {
			words = append(words, word)
			word = ""
		}
	}
	fmt.Println(words)
}

func main() {
	args := os.Args[1:]
	RoString(args[0])
}
