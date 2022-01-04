package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after call", myString)
}

func changeUsingPointer(s *string) {
	newValue := "red"
	*s = newValue
}
