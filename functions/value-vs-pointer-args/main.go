package main

import (
	"fmt"
)

func main() {
	leedsDockEmployees := map[string]int{
		"Pikachu":       123456,
		"Squidward Dab": 234561,
		"Rick Sanchez":  345612,
		"Gandalf":       456123,
		"Jon Snow":      561234,
		"Skywalker":     612345,
	}

	name := "Pikachu"
	if id, err := getSkyID(name, leedsDockEmployees); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sky ID of employee %v is %v \n\n", name, id)
	}
}

// value vs pointer param, mainly depending on if you need to mutate the value of argument
// pointer could help with performance if you pass a large data structure
// reference types will always be passed as pointers
func getSkyID(name string, employees map[string]int) (int, error) {
	id, ok := employees[name]
	if ok != true {
		return 0, fmt.Errorf("The employee %v doesn't exist", name)
	}

	// uncomment line below, name will not changed to "wei wei" unless we change name into type *string
	// name = "wei wei"

	// uncomment line below, id will be changed to 999999 as map is reference type (internal pointer)
	// id = 999999

	return id, nil
}
