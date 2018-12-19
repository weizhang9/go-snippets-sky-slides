package main

import "fmt"

type aStruct struct{ name string }

func main() {
	var nilPointer *aStruct
	fmt.Println(nilPointer) // passing a nil pointer in your program will trigger runtime panic

	var zeroValuePointer *aStruct = new(aStruct)
	fmt.Println(zeroValuePointer) // initialised a zero-valued aStruct pointer to be assigned new value
	// since zeroValuePointer is a pointer type, compiler will insert * in front of it
	// which means no need to do (*zeroValuePointer).name = "Gopher"
	zeroValuePointer.name = "Gopher"
	fmt.Println(zeroValuePointer)
	fmt.Printf("\nVariable zeroValuePointer's address in memory is %v\n\n", &zeroValuePointer)

	// pointer arithmetic is not supported in Go, but if you have to, package unsafe will allow you to do that
	// https://golang.org/pkg/unsafe/
	is := []int{1, 2, 3}
	baseAdd := &is[0]
	secAdd := &is[1] // address is exactly 4 bit ahead of baseAdd, but we can't add 4 on baseAdd to get here like in C
	fmt.Printf("value of is: %v\nAddress of baseAdd is: %p\nAddress of secAdd is: %p", is, baseAdd, secAdd)
}
