package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if (i == 1 && j == 1) || (i == y && j == x) {
					fmt.Print("/")
				} else if (i == y && j == 1) || (i == 1 && j == x) {
					fmt.Print("\\")
				} else if (j == 1) || (j == x) || (i == 1) || (i == y) {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
