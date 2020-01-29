package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		count := 0
		num := 0
		num, _ = strconv.Atoi(args[0])
		l := PrimeFactors(num)

		for i := 0; i < len(l); i++ {
			count++
			if count >= i && i > 0 {
				fmt.Print("*")
			}
			fmt.Print(l[i])
		}
		fmt.Println()

	} else {
		fmt.Println()
	}
}

func PrimeFactors(x int) []rune {
	factors := []rune{}

	candidate := 2
	for candidate <= x {
		if x%candidate == 0 {
			factors = append(factors, rune(candidate))
			x = x / candidate
		} else {
			candidate++
		}
	}

	return factors
}
