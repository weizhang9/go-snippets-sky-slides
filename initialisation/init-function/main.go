package main

import (
	"fmt"
)

func init() {
	fmt.Println("init() doesn't take argument nor return any value, it's used to verify or repair correctness of the program state before real execution begins")
}

func main() {
	fmt.Println("Now the program is in prime condition to run tasks.")
}
