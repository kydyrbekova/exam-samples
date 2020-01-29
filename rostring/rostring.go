package main

import (
	"os"
	"strings"

	"fmt"
)

func main() {

	new := ""
	new2 := ""
	if len(os.Args) > 1 && len(os.Args) <= 2 {
		str := os.Args[1]
		for i := 0; i < len(str); i++ {
			str = standardizeSpaces(str)
			if str[0] != 32 {
				for v := 0; v < len(str); v++ {

					if str[v] == 32 {
						new = str[:v]
						new2 = str[v:]
						str = new2 + " " + new
						fmt.Println(standardizeSpaces(str))
						return
					}
				}
			} /*else if str[0]==32{
				fmt.Println(standardizeSpaces(str))
					return
			}*/
		}
	} else {
		fmt.Println()
	}
}
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func Join(strs []string, sep string) string {
	strV := ""

	for i := 0; i < len(strs); i++ {
		strV += strs[i]
		if i < len(strs)-1 {
			strV += sep
		}
	}
	return strV
}
