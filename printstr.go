package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	r1 := []rune(s)
	for _, ch := range r1 {
		z01.PrintRune(ch)
	}
}
