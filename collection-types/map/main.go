package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map declaration syntax: map[key-type]value-type")
	// can initialise an empty map and add elements later (e.g. add elements from loops)
	squads := make(map[[3]string]int)
	squads = map[[3]string]int{
		{"cosmonauts", "news-and-sports-ng", "UK"}: 1097728,
		{"hoff", "sports-ng", "DE"}:                1099776,
	}
	fmt.Printf("\n%v\n\n", squads)

	// the value of an uninitialized map is nil
	// can NOT add element into nil map, otherwise runtime panic trigger
	var nilMap map[int]bool
	//	nilMap[0] = false
	fmt.Printf("nilMap's type is %T, value is %v\n\n", nilMap, nilMap)

	// a map with key type of string and value type of int
	// key type can NOT be of a function, map, or slice
	timeZone := map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	// add an element into a map
	timeZone["GMT"] = 0 * 60 * 60
	fmt.Println(timeZone) // map is unordered, order of returned elements may be different from how you initilised

	// delete an element from a map
	delete(timeZone, "UTC")
	fmt.Printf("%v\n\n", timeZone)
	
	// query an element's key after deletion of the element is useless
	fmt.Println(timeZone["UTC"])
	fmt.Println(timeZone["WTF"])
	
	// use built-in comma ok syntax to check if an element exists
	_, ok := timeZone["UTC"] // could throw away the element value with blank identifier
	fmt.Println(ok)
}
