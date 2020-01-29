package main

import "github.com/01-edu/z01"
import "os"


func main() {
	if len(os.Args) == 3 {
		text := os.Args[1] + os.Args[2]
		ftext := ""
		check := true

		for _, t := range text {
			check = true
			for _, f := range ftext {
				if t == f {
					check = false
				}
			}
			if check == true {
				ftext = ftext + string(t)
			}
		}
		for _, l := range ftext {
			z01.PrintRune(rune(l))
		}
	}
	z01.PrintRune(10)
}
