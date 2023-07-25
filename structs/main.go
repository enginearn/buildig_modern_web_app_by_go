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
}
