package main

import (
	"fmt"
)

func main() {
	for i := 9; i > 0; i-- {
		fmt.Println(i)
	}

	// can use multi-assign to utilise 2 variables
	for i, j := 0, 9; i < 10; i, j = i+1, j-1 {
		fmt.Printf("%v\t%v\n", i, j)
	}

	// while loop
	k := 0
	for k < 3 {
		fmt.Println(k)
		k++
	}

	// use label to break off outer loop
BreakItOff:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i * j)
			if i*j >= 5 {
				break BreakItOff
			}

		}
	}

	// for range loop - string, array/slice/map, channel
	s := "merry christmas"
	for k, v := range s {
		fmt.Printf("The letter is: %v \t Key of the letter: %v \t UTF8 code point of the letter: %v \t\n", string(v), k, v)
	}
	a := [...]string{"s", "m", "i", "l", "e"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
