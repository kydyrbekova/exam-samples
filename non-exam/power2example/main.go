package main

import "fmt"

func main() {
	nbr := 5
	ans := power(nbr)
	fmt.Println(ans)
}

func power(nbr int) bool {
	if nbr < 1 {
		return false
	}
	for nbr > 1 {
		if nbr%2 == 0 {
			nbr = nbr/2
		} else {
			return false
		}
	}
	return true
}