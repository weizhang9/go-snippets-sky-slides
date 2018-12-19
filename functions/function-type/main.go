package main

import (
	"fmt"
)

func main() {
	// beware when function assigned to a variable, it won't be hoistered
	var divide func(float64, float64) (float64, error) // uninitialised function type has a value of nil

	divide = func(x, y float64) (res float64, err error) {
		if y == 0.0 {
			return 0.0, fmt.Errorf("Check out Riemann Sphere if you really want to divide by zero")
		}
		res = x / y
		return
	}

	dvs, err := divide(17.0, 1.888)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dvs)
}
