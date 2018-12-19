package main

import "fmt"

func main() {
	// if you don't set the capacity, it will be the length of your slice
	// can custom set index on values like with array
	alphabet := []string{"a", 4: "b", "c"}
	fmt.Println("Slice alphabet", alphabet, "has length of", len(alphabet), "with capacity of", cap(alphabet))

	// slices hold reference, modifying item value will mutate underlying array's value
	aClone := alphabet
	aClone[1] = "second-martini"

	fmt.Println(alphabet)
	fmt.Println(aClone)

	// ways to resize slices
	fmt.Printf("\nSlicing from here:\n")
	sliceAllItems := alphabet[:]
	sliceFromCertainItem := alphabet[2:]
	sliceUpToCertainItem := alphabet[:4] // [starting-item:ending-item) - ending item is not included
	sliceFromCertainItemUpToCetainItem := alphabet[3:5]
	fmt.Println(sliceAllItems)
	fmt.Println(sliceFromCertainItem)
	fmt.Println(sliceUpToCertainItem)
	fmt.Println(sliceFromCertainItemUpToCetainItem)

	// use append() to resize slices, can append multiple SAME TYPE values
	alphabet = append(alphabet, "one-more-drink", "sorry-two-more-drinks")
	// BUT if you have to append a slice, use ... to decomposite it
	alphabet = append(alphabet, []string{"now-pub-crawl", "bye"}...)
	fmt.Println("Slice alphabet is now", alphabet, "has length of", len(alphabet), "with capacity of", cap(alphabet))

	// the moment you sliced an array, values of the array will be changed accordingly when you mutate the values of the slice
	fmt.Printf("\nSlicing an array could have unexpected result:\n")
	array := [5]int{6, 7, 8, 9, 10}
	aSlice := array[0:4]
	aSlice[2] = 2222
	fmt.Println(aSlice)
	fmt.Println(array)
}
