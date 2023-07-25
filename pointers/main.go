package main

import "log"

func main() {
	var name string = "John"
	var namePointer *string = &name

	log.Println("Name: ", name)
	log.Println("NamePointer: ", namePointer)
	log.Println("NamePointer value: ", *namePointer)

	name = "Doe"
	log.Println("Name: ", name)
	log.Println("NamePointer: ", namePointer)
	log.Println("NamePointer value: ", *namePointer)

	*namePointer = "Jane"
	log.Println("Name: ", name)
	log.Println("NamePointer: ", namePointer)
	log.Println("NamePointer value: ", *namePointer)

	changeUsingPointer(&name)
	log.Println("Name: ", name)
}

func changeUsingPointer(namePointer *string) {
	*namePointer = "太郎"
}
