package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for firstDigit := '0'; firstDigit <= '7'; firstDigit++ {
		for secondDigit := '1'; secondDigit <= '8'; secondDigit++ {
			if secondDigit > firstDigit {
				for thirdDigit := '2'; thirdDigit <= '9'; thirdDigit++ {
					if thirdDigit > secondDigit {
						z01.PrintRune(firstDigit)
						z01.PrintRune(secondDigit)
						z01.PrintRune(thirdDigit)
						if firstDigit == '7' && secondDigit == '8' && thirdDigit == '9' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
