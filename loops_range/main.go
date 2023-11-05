package main

import (
	// "fmt"
	"log"
)

func main() {
	for i := 0; i <= 2; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "house", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	animals_map := make(map[string]string)
	animals_map["dog"] = "太郎"
	animals_map["cat"] = "おかゆ"

	for animalType, animal := range animals_map {
		log.Println(animalType, animal)
	}

	firstLine := "Once upon a midnight dreary"

	for i, w := range firstLine {
		log.Println(i, ":", w)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 31})
	users = append(users, User{"Mary", "Smith", "mary@smith.com", 28})
	users = append(users, User{"Anny", "Moris", "anny@moris.com", 19})
	users = append(users, User{"Marin", "Houshou", "main@captain.com", 17})

	for i, u := range users {
		log.Println(i, u.FirstName, u.LastName, u.Email, u.Age)
	}
}
