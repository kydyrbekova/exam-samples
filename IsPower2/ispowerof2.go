package main

import "os"
import "strconv"
import "github.com/01-edu/z01"

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {

		our_int, err := strconv.Atoi(arg[0])
		if err != nil {
			for _, e := range string(err.Error()) {
				z01.PrintRune(e)
			}
		} else {

			if IsPowerof2(our_int) {
				z01.PrintRune('t')
				z01.PrintRune('r')
				z01.PrintRune('u')
				z01.PrintRune('e')

			} else {
				z01.PrintRune('f')
				z01.PrintRune('a')
				z01.PrintRune('l')
				z01.PrintRune('s')
				z01.PrintRune('e')
			}
		}
		z01.PrintRune(10)
	} else {

		z01.PrintRune(10)
	}
}


func IsPowerof2(num int) bool {
	if num < 1 {
		return false
	}
	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			return false
		}

	}
	return true
}
