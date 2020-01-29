package main

import "fmt"

func main() {
	fmt.Println(byte('A'))
	fmt.Println(SwapBits(byte('A')))
}

func SwapBits(octet byte) byte {
	return ((octet&0x0F)<<4 | (octet&0xF0)>>4)
}
