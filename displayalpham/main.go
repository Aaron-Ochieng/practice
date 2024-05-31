package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arr := [][]rune{}
	var ax []rune

	for i := 'a'; i <= 'z'; i += 3 {
		if i+3 <= 'z' {
			temp := i
			for x := i; x <= temp+2; x++ {
				ax = append(ax, x)
			}
			arr = append(arr, ax)
			ax = []rune{}

		} else {
			for x := i; x <= 'z'; x++ {
				ax = append(ax, x)
			}
			arr = append(arr, ax)
		}
	}

	for i := 0; i < len(arr); i++ {
		if (i+1)%2 == 0 {
			for _, rn := range arr[i] {
				z01.PrintRune(rn - 32)
			}
		} else {
			for _, rn := range arr[i] {
				z01.PrintRune(rn)
			}
		}
	}
	z01.PrintRune('\n')
}
