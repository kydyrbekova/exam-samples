package main

import "fmt"

//Tester func
func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

// Your func
func Compare(a, b string) int {
	if a < b {
		return -1
	} else
	if a > b {
		return 1
	}
	return 0
}