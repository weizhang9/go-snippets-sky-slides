package main

import (
	"fmt"
)

func main() {
	a := 19 // 19 = 10011 in binary
	fmt.Printf("%b\n", a)

	b := 17 // 17 = 10001 in binary
	fmt.Printf("%b\n", b)

	fmt.Println(a & b)  // & includes what both bits are set - 10011 & 10001 = 10001, which is 17 in decimal
	fmt.Println(a | b)  // | includes what either bit is set - 10011 | 10001 = 10011, which is 19 in decimal
	fmt.Println(a ^ b)  // ^ includes what one or the other bit is set, but not both - 10011 ^ 10001 = 00010, which is 2 in decimal
	fmt.Println(a &^ b) // &^ use & on ^b - 10011 &^ 10001 = 10011 & 01110 = 00010, which is 2 in decimal
}
