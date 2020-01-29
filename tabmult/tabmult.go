package main

import "github.com/01-edu/z01"
import "strconv"


func main() {
	//fake arguments
	osArgs := []string{"sadsdsas","1444411"}
	l := 0
	for range osArgs {
		l++
	}
	arg1 := osArgs[1]
	if l == 2 {
		arg, _ := strconv.Atoi(osArgs[1])

		for i := 1; i <= 9; i++ {
			ans := intstr(i*arg)
			z01.PrintRune(rune(i)+48)
			z01.PrintRune(' ')
			z01.PrintRune('x')
			z01.PrintRune(' ')
			for a, _ := range arg1 {
				z01.PrintRune(rune(arg1[a]))
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			for b, _ := range ans {
				z01.PrintRune(rune(ans[b]))
			}
			z01.PrintRune(10)
		}
	}
}

func intstr(n int) string {
	toReturn := ""
	for n != 0 {
		toReturn = string(rune(n%10)+'0') + toReturn
		n /= 10
	}
	return toReturn
}