package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	array := os.Args[1:]
	count := 0
	for range array {
		count++
	}
	a := make([]string, count)
	a = array
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if a[i] > a[j] {
				k := a[i]
				a[i] = a[j]
				a[j] = k
			}
		}

	}
	for i := 0; i < count; i++ {
		run_ := []rune(a[i])
		for _, value := range run_ {
			z01.PrintRune(value)
		}
	}

	z01.PrintRune(10)
}
