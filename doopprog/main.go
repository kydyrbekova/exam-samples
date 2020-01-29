package main

import "github.com/01-edu/z01"
import "strconv"

func main() {
	//fake arguments
	//osArgs := []string{"0", "5", "+", "2"}
	//osArgs := []string{"0", "hello", "+", "1"}
	//osArgs := []string{"0", "5", "a", "1"}
	osArgs := []string{"0", "1", "/", "-1"}
	//osArgs := []string{"0", "5", "/", "0"}
	//osArgs := []string{"0", "5", "%", "2"}
	//osArgs := []string{"0", "5", "%", "0"}
	//osArgs := []string{"0", "5", "*", "2"}

	l := 0
	for range osArgs {
		l++
	}
	if l == 4 {
		arg1 := osArgs[1]
		arg2 := osArgs[2]
		arg3 := osArgs[3]

		var ans int
		var error string
		solveornot := false

		arg1int, _ := strconv.Atoi(arg1)
		arg3int, _ := strconv.Atoi(arg3)

		arg1rune := []rune(arg1)
		arg3rune := []rune(arg3)

		arg1nbrs := 0
		arg3nbrs := 0
		if arg2 == "+" || arg2 == "/" || arg2 == "%" || arg2 == "*" {
			solveornot = true
		}

		for i := 0; i < len(arg1rune); i++ {
			if arg1rune[i] >= '0' && arg1rune[i] <= '9' {
				arg1nbrs++
			}
		}

		for i := 0; i < len(arg3rune); i++ {
			if arg3rune[i] >= '0' && arg1rune[i] <= '9' {
				arg3nbrs++
			}
		}
		if arg1nbrs == len(arg1rune) && arg3nbrs == len(arg3) && solveornot == true {
			if arg2 == "+" {
				ans = arg1int + arg3int
			} else if arg2 == "/" {
				if arg3int != 0 {
					ans = arg1int / arg3int
				} else {
					error = "No divison by 0"
				}
			} else if arg2 == "%" {
				if arg3int != 0 {
					ans = arg1int % arg3int
				} else {
					error = "No modulo by 0"
				}
			} else if arg2 == "*" {
				ans = arg1int * arg3int
			} else {
				ans = 0
			}
			ans_str := intstr(ans)

			if error == "" {
				for a, _ := range ans_str {
					z01.PrintRune(rune(ans_str[a]))
				}
			} else {
				for a, _ := range error {
					z01.PrintRune(rune(error[a]))
				}
			}
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune(10)
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
