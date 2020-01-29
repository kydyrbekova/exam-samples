package main

import "github.com/01-edu/z01"

func main() {
	//fake arguments
	osArgs := []string{"sadsdsas", "Hello There", "sdsdsdsd"}

	l := 0
	for range osArgs {
		l++
	}
	if l > 1 {
		arg_rune := []rune(osArgs[1])
		l_arg := 0
		for range osArgs[1] {
			l_arg++
		}

		for i := 0; i < l_arg; i++ {
			if arg_rune[i] >= 'A' && arg_rune[i] < 'M' {
				arg_rune[i] = arg_rune[i]+13
			} else if arg_rune[i] > 'M' && arg_rune[i] <= 'Z' {
				arg_rune[i] = arg_rune[i]-13
			} else if arg_rune[i] >= 'a' && arg_rune[i] < 'm' {
				arg_rune[i] = arg_rune[i]+13
			} else if arg_rune[i] > 'm' && arg_rune[i] <= 'z' {
				arg_rune[i] = arg_rune[i]-13
			}
			z01.PrintRune(arg_rune[i])
		}
		z01.PrintRune(10)
	}
}
