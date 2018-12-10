package main

import (
	"fmt"
)

func main() {
	// boolean - can be used for logical tests
	a := 1+1 == 2
	b := 1+1 == 3
	fmt.Printf("value: %v\t type: %T\n", a, a)
	fmt.Printf("value: %v\t type: %T\n", b, b)
	var c bool
	fmt.Printf("value: %v\t type: %T\n", c, c)

	// numberic types
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	// string
	s := "my lovely string "
	fmt.Printf("value: %v\t type: %T\n", s[5], s[5])
	// string is the only immutable primitive, but can be concatenated
	s1 := "now baby string wants to join the A team"
	fmt.Println(s + s1)
	// string can be converted into a slice of bytes/runes as showed in constants
}
