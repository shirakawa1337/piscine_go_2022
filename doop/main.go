package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "+" {
			if Convert(args[i-1]) == -99 || Convert(args[i+1]) == -99 {
				return
			}
			res, ch := Plus(Convert(args[i-1]), Convert(args[i+1]))
			if ch {
				fmt.Print(res)
				fmt.Print("\n")
			} else {
				return
			}
		} else if args[i] == "-" {
			if Convert(args[i-1]) == -99 || Convert(args[i+1]) == -99 {
				return
			}
			res, ch := Minus(Convert(args[i-1]), Convert(args[i+1]))
			if ch {
				fmt.Print(res)
				fmt.Print("\n")
			} else {
				return
			}
		} else if args[i] == "*" {
			if Convert(args[i-1]) == -99 || Convert(args[i+1]) == -99 {
				return
			}
			res, ch := Mult(Convert(args[i-1]), Convert(args[i+1]))
			if ch {
				fmt.Print(res)
				fmt.Print("\n")
			} else {
				return
			}
		} else if args[i] == "/" {
			if Convert(args[i-1]) == -99 || Convert(args[i+1]) == -99 {
				return
			}
			res, ch := Devide(Convert(args[i-1]), Convert(args[i+1]))
			if ch {
				fmt.Print(res)
				fmt.Print("\n")
			} else {
				fmt.Print("No division by 0")
				fmt.Print("\n")
			}
		} else if args[i] == "%" {
			if Convert(args[i-1]) == -99 || Convert(args[i+1]) == -99 {
				return
			}
			res, ch := Modulo(Convert(args[i-1]), Convert(args[i+1]))
			if ch {
				fmt.Print(res)
				fmt.Print("\n")
			} else {
				fmt.Print("No modulo by 0")
				fmt.Print("\n")
			}
		}
	}
}

func Compare(a, b int) (int, int) {
	max := 0
	min := 0
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return max, min
}

func Plus(a, b int) (int, bool) {
	max, min := Compare(a, b)
	res := a + b
	if max >= 0 && min >= 0 {
		if max > res {
			return 0, false
		}
	} else if max < 0 && min < 0 {
		if max < res {
			return 0, false
		}
	}
	return res, true
}

func Minus(a, b int) (int, bool) {
	max, min := Compare(a, b)
	res := a - b
	if max >= 0 && min >= 0 {
		if max < res && res > 0 {
			return 0, false
		}
	} else if max > 0 && min < 0 {
		if max < res {
			return 0, false
		}
	} else if max < 0 && min > 0 {
		if max > res {
			return 0, false
		}
	}
	return res, true
}

func Mult(a, b int) (int, bool) {
	max, min := Compare(a, b)
	res := a * b

	if max >= 0 && min >= 0 {
		if max > res {
			return 0, false
		}
	} else if max < 0 && min < 0 {
		if max < res {
			return 0, false
		}
	}
	return res, true
}

func Devide(a, b int) (int, bool) {
	res := 0
	if b == 0 {
		return 0, false
	} else {
		res = a / b
	}
	return res, true
}

func Modulo(a, b int) (int, bool) {
	res := 0
	if b == 0 {
		return 0, false
	} else {
		res = a % b
	}
	return res, true
}

func Convert(arg string) int {
	num := 0
	change := true
	for i := 0; i < len(arg); i++ {
		if arg[i] == '-' {
			change = false
			i++
		} else if !(arg[i] >= '0' && arg[i] <= '9') {
			return -99
		}

		temp := num
		num = num*10 + int(arg[i]) - '0' ///
		if temp > num && temp > 0 {
			return -99
		} else if temp > num && temp < 0 {
			return -99
		}

	}
	if !change {
		num = num * -1
	}

	return num
}
