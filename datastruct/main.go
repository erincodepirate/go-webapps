package main

import "log"

type myStruct struct {
	FirstName string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	otherMap := make(map[string]myStruct)

	myUser := myStruct{FirstName: "Kirby"}

	otherMap["foo"] = myUser

	log.Println(otherMap["foo"].FirstName)

	var mySlice []string

	mySlice = append(mySlice, "foo")
	mySlice = append(mySlice, "bar")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])
}
