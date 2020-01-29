package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!"
	nb := FirstRune(str)
	z01.PrintRune(nb)
	z01.PrintRune(10)
}

// Your code
func FirstRune(str string) rune{
	str_rune := []rune(str)
	return str_rune[0]
}