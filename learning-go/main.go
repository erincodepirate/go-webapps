package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var what2Say string
	what2Say = "What"

	fmt.Println("My name is " + what2Say + "!")

	var i int
	i = 420
	fmt.Println(i, "is my number")

	whatWasSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid)
}

func saySomething() string {
	return "something"
}
