package main

import "log"

func main() {
	cat := "cat"

	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is not true")
	} else if myNum < 100 && isTrue {
		log.Println("blah")
	} else {
		log.Println("blah")
	}

	switch cat {
	case "cat":
		log.Println("cat is cat")
	case "dog":
		log.Println("cat is dog")
	case "fish":
		log.Println("cat is fish")
	default:
		log.Println("cat is something")
	}
}
