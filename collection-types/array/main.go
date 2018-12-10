package main

import (
	"fmt"
)

func main() {
	// can assign index to value, but it must be constant integer EXPRESSION, and cannot be out of range
	// when assign a higher index than its postion to an item, the missing items will be initialised as the default value of the item type, i.e. 0 for int, 0.0 for float, empty for string etc
	// use ... to fit the size of the length, useful for when not knowing the array size when declaring, or just want to save memory and improve performance
	alphabet := [...]string{"a", "b", 2 << 1: "c", "d"}  // 2 << 1 = 4; or you can do 2*2 or 2+2
	for i := range alphabet {
		fmt.Println("Index", i, "is", alphabet[i])
	}
	
	// only passed a copy to aClone
	aClone := alphabet
	aClone[1] = "martini"
	fmt.Println(alphabet)
	fmt.Println(aClone)
	
	// assign pointer of original array to mutate original value(s)
	aCloneNoQuit := &alphabet
	aCloneNoQuit[1] = "i-just-want-some-martini"
	fmt.Println(alphabet)
	fmt.Println(aCloneNoQuit)	
}
