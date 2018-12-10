package main

import (
	"fmt"
)

func main() {
	// use defer to postpone execution of a statement/function until the end of the calling function
	doDBOperations()
}

func connectToDB() {
	fmt.Println("ok, connected to database")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from database")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("defering the database disconnect")
	defer disconnectFromDB()
	fmt.Println("doing some database operations")
	fmt.Println("sorry, I crashed")
	fmt.Println("ok, I am back on track")
	return //terminate the program
	// deferred function executed here just before actually returning, even if there is a return or abnormal termination (panic) before
}
