package main

import "fmt"

func main() {
	defer rescue()
	fmt.Println("Time to play with death, fasten your seatbelt a little bit.")
	fmt.Printf("\nRecover is only useful when called inside a deferred function as it retreive a non-nil error value\nIn normal execution a call to recover will return nil and have no other effect.\n\n")
	fmt.Println("Rule of thumb:\n1) Always recover from panic in your package\n2) Return errors as error values to the callers of your package")
	badCall()

	fmt.Println("I never get executed as program stopped at panic(), then recover() takes the error value and moved on.")
}

func badCall() {
	panic("oh no, I'm dying!")
}

func rescue() {
	if err := recover(); err != nil {
		fmt.Printf("\nBabylon is panicking again - \"%s\" 🙄🙄🙄\n", err)

	}
}
