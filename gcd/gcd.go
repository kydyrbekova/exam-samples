package main

import "fmt"
import "os"
import "strconv"

func main() {
	if len(os.Args) == 3 {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[2])
		lower := a
		upper := b
		if a > b {
			lower = b
			upper = a
		}
		for i := lower; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				fmt.Print(i)
				break
			}
		}
		fmt.Println()
	}
}
