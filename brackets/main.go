package main

import (
	"fmt"
	"os"
)

func correctlyFormatted(s string) {
	extractedbrackets := ""

	stack := []rune{}
	for _, char := range s {
		if char == '[' || char == ']' || char == '(' || char == ')' || char == '{' || char == '}' {
			extractedbrackets += string(char)
		}
	}
	fmt.Println(extractedbrackets)

	for _, char := range extractedbrackets {
		if char == '[' || char == '{' || char == '(' {
			stack = append(stack, char)
		} else if len(stack) > 0 {
			if (char == ']' && stack[len(stack)-1] == '[') || (char == '}' && stack[len(stack)-1] == '{') || (char == ')' && stack[len(stack)-1] == '(') {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
	}
	if len(stack) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("Error")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error")
	}
	correctlyFormatted(args[0])
}
