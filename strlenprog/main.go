package main

import "fmt"

func main() {
	str := "Hello World!"
	nb := Strlen(str)
	fmt.Println(nb)
}

// Your code
func Strlen(str string) int{
	l := 0
	for range str {
		l++
	}
	return l
}