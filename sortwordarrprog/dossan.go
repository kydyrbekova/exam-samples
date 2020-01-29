package main

import (
	"fmt"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(array []string) {
	l := len(array)
	for j := 0; j < l-1; j++ {
		for i := 0; i < l-1-j; i++ {
			cur := array[i]
			next := array[i+1]
			if rune(cur[0]) > rune(next[0]) {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}
}
