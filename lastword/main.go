package main

import "github.com/01-edu/z01"

func main() {
	//fake arguments
	osArgs := []string{"0", " dsvf fdgfasfdf ggg ggg    "}

	l := 0
	for range osArgs {
		l++
	}
	if l == 2 {
		arg_rune := []rune(osArgs[1])
		l_arg := 0
		for range osArgs[1] {
			l_arg++
		}

		for i := l_arg-1; i > 0; i-- {
			if arg_rune[i] != ' ' {
				for j := i; j > 0; j-- {
					if arg_rune[j] == ' ' {
						for k := j+1; k <= i; k++ {
							z01.PrintRune(arg_rune[k])
						}
						break
					}
				}
				break
			}
		}
	}
	z01.PrintRune(10)
}
