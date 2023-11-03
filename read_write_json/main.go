package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"Name"` // `json:"Name"` is called a tag
	Age  int    `json:"Age"`  // `json:"Age"` is called a tag
	Car  string `json:"Car"`  // `json:"Car"` is called a tag
}

func main() {
	Json := `[
		{
			"Name": "John",
			"Age": 30,
			"Car": null
		},
		{
			"Name": "Jane",
			"Age": 25,
			"Car": "Ford"
		}
	]`

	fmt.Println(Json)

	var people []Person
	err := json.Unmarshal([]byte(Json), &people)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(people)

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i].Name)
		fmt.Println(people[i].Age)
		fmt.Println(people[i].Car)
	}

	// write json from struct
	people2 := []Person{
		{Name: "John", Age: 30, Car: "Hpnda"},
		{Name: "Jane", Age: 25, Car: "Ford"},
		{Name: "Jack", Age: 20, Car: "Toyota"},
		{Name: "Jill", Age: 15, Car: "Nissan"},
		{Name: "Joe", Age: 10, Car: "GMT"},
	}

	Json2, err := json.Marshal(people2)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(Json2))

	var mySlice []Person

	var m1 Person
	m1.Name = "John"
	m1.Age = 30
	m1.Car = "Honda"

	var m2 Person
	m2.Name = "Jane"
	m2.Age = 25
	m2.Car = "Ford"

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	mewJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshelling ", err)
	}

	log.Println(string(mewJson))

	fmt.Println(mySlice)
}
