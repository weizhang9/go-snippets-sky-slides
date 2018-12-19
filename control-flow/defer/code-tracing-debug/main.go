package main

import "fmt"

func main() {
	skyNews()
}

func trace(s string) string {
	fmt.Println("entering: ", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving: ", s)
}

func skyNews() {
	defer untrace(trace("Sky News"))
	fmt.Println("in Sky News")
	skySports()
}

func skySports() {
	defer untrace(trace("Sky Sports"))
	fmt.Println("in Sky Sports")
}
