package main

import (
	"fmt"
	"os"
	"strconv"
)

func TabMult(n int) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d"+"x"+"%d"+"="+"%d\n", i, n, (i * n))
	}
}

func main() {
	args := os.Args[1:]
	if len(args)-1 > 0 || len(args)-1 < 0 {
		return
	}
	y, _ := strconv.Atoi(args[0])
	TabMult(y)
}
