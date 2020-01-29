package main

import "fmt"

func main() {
	nbr := 1337
	str := itoa(nbr)
	fmt.Println(str)
}

func itoa(nbr int) string {
	res := ""
	for nbr != 0 {
		res = string(rune(nbr%10+'0')) + res
		nbr = nbr / 10
	}
	return res
}