package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for index, value := range arguments[0] {
		if index > 1 {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}
