package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	for i := 'z'; i <= 'z' && i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	fmt.Print("\n")
}
