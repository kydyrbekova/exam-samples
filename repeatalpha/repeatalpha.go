package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	len_args := 0
	for range args {
		len_args++
	}

	if len_args == 1 {
		str := args[0]
		new_str := ""
		alpha := "abcdefghijklmnopqrstuvwxyz"
		Alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for i := 0; i < len(str); i++ {
			if str[i] > 'a' && str[i] < 'z' {
				for j := 0; j < len(alpha); j++ {
					if string(str[i]) == string(alpha[j]) {
						for k := 0; k < j+1; k++ {
							new_str = new_str + string(str[i])
						}
					}
				}

			} else if str[i] > 'A' && str[i] < 'Z' {
				for j := 0; j < len(Alpha); j++ {
					if string(str[i]) == string(Alpha[j]) {
						for m := 0; m < j+1; m++ {
							new_str = new_str + string(str[i])
						}
					}
				}
			} else {
				new_str = new_str + string(str[i])
			}

		}
		fmt.Println(new_str)

	} else {
		z01.PrintRune(10)
	}
}
