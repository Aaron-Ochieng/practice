package main

import (
	"fmt"
	"os"
)

func RemoveAt(str, sep string) string {
	word:=""
	for i := 0; i < len(str); i++ {
		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {
			word+=" "
			i += len(sep) - 1
		} else {
			word+=string(str[i])
		}
	}
	fmt.Println(word)
	return word
}


func main() {
	args := os.Args[1:]
    fmt.Println(RemoveAt(args[0], args[1]))
}