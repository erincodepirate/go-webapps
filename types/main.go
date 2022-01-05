package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Trevor",
		LastName:  "Belmont",
	}

	log.Println(user.FirstName, user.LastName)
}
