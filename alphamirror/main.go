package main

import "github.com/01-edu/z01"

func main() {
	//fake arguments
	osArgs := []string{"0", "My Horse Is Amazing."}

	l := 0
	for range osArgs {
		l++
	}
	if l == 2 {
		arg1 := []rune(osArgs[1])
		larg1 := 0
		for range arg1 {
			larg1++
		}
		for i := 0; i < larg1; i++ {
			if arg1[i] >= 65 && arg1[i] <= 90 {
				arg1[i] = 90-(arg1[i]-65)
			} else
			if arg1[i] >= 97 && arg1[i] <= 122 {
				arg1[i] = 122-(arg1[i]-97)
			}
			z01.PrintRune(rune(arg1[i]))
		}
	}
	z01.PrintRune(10)
}
