package main

import "github.com/01-edu/z01"

func displayalrevm() {
	counter := 1

	for i := 'z'; i >= 'a'; i-- {
		if i%2 != 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
		counter++
	}
}

func main() {
	displayalrevm()
	z01.PrintRune('\n')
}
