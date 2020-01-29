package main

import (
	"fmt"
	"os"
)

var err int

func check(s string) bool {

	for _, c := range s {
		if c == ' ' || c < '0' || c > '9' {
			return false
		}
	}
	return true
}
func Atoi(str string) int {
	x := 0
	var this string
	ok := true
	change := false
	second := false
	for _, c := range str {
		if c != '+' && c != '-' {
			ok = false
		}
		if ok == true {
			if second == true {
				return 0
			}
			if c == '-' {
				change = true
			}
			second = true
		} else {
			this = this + string(c)
		}
	}
	if check(this) == true {
		for _, c := range this {

			cnt := 0
			for i := '1'; i <= c; i++ {
				cnt++
			}
			x = x*10 + cnt
		}
	} else {
		err = 1
	}
	if change == true {
		x *= -1
	}
	return x
}

func main() {
	arg := os.Args[1:]
	ln := 0
	for i := range arg {
		ln = i + 1
	}
	if ln == 3 {
		a := Atoi(arg[0])
		b := Atoi(arg[2])
		if arg[1] != "*" && arg[1] != "/" && arg[1] != "+" && arg[1] != "-" && arg[1] != "%" || err == 1 {
			fmt.Println(0)
		} else {
			if b == 0 && arg[1] == "/" {
				fmt.Println("No division by 0")
			} else if b == 0 && arg[1] == "%" {
				fmt.Println("No Modulo by 0")
			} else {
				c := 0
				if arg[1] == "*" {
					c = a * b
				}
				if arg[1] == "/" {
					c = a / b
				}
				if arg[1] == "+" {
					c = a + b
				}
				if arg[1] == "-" {
					c = a - b
				}
				if arg[1] == "%" {
					c = a % b
				}
				fmt.Println(c)
			}
		}
	}
}
