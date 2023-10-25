package main

import (
	"log"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var s = "seven"

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething fnc is", s)
	return s, "world"
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	user1 := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "sample@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Println("user1: ", user1)
	log.Println("user1 ID: ", user1.ID)
	log.Println("user1 First Name: ", user1.FirstName)
	log.Println("user1 Last Name: ", user1.LastName)
	log.Println("user1 Email: ", user1.Email)
	log.Println("user1 Created At: ", user1.CreatedAt)
	log.Println("user1 Updated At: ", user1.UpdatedAt)

	user2 := User{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane_doe@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Println("user2: ", user2)
	log.Println("user2 ID: ", user2.ID)
	log.Println("user2 First Name: ", user2.FirstName)
	log.Println("user2 Last Name: ", user2.LastName)
	log.Println("user2 Email: ", user2.Email)
	log.Println("user2 Created At: ", user2.CreatedAt)
	log.Println("user2 Updated At: ", user2.UpdatedAt)

	var s2 = "six"

	s := "five"

	log.Println("s: ", s)
	log.Println("s2: ", s2)

	saySomething("hello")

	var ms myStruct
	ms.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Jane",
	}

	log.Println("ms: ", ms.printFirstName())
	log.Println("ymVar2: ", myVar2.printFirstName())
}
