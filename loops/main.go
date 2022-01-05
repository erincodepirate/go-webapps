package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	anmap := make(map[string]string)
	anmap["dog"] = "Snoop"
	anmap["cat"] = "Garfield"

	for animalType, animal := range anmap {
		log.Println(animalType, animal)
	}

	str := "foo bar"

	for i, l := range str {
		log.Println(i, ":", l)
	}

}
