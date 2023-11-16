package main

import "fmt"

// import "github.com/01-edu/z01"

// func main() {
// 	counter := '0'
// 	for i := 0; i < 50; i++ {
// 		counter++
// 	}
// 	z01.PrintRune(rune(counter))
// }

func main() {
	s := "HowHAareHAyou"
	sep := "HA"
	runes := []rune(s)
	fmt.Println(string(runes[3 : 3+len(sep)]))
}

// HamzaKaBoom,KaB
