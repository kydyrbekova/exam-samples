package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!"
	nb := LastRune(str)
	z01.PrintRune(nb)
	z01.PrintRune(10)
}

// Your code
func LastRune(str string) rune{
	l := 0
	for range str {
		l++
	}
	str_rune := []rune(str)
	return str_rune[l-1]
}