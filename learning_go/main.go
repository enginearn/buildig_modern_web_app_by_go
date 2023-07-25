package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var whatToSay string = "Goodbye cruel world!"
	fmt.Println(whatToSay)

	var i int = 7
	fmt.Println("i is set to ", i)

	var f float32 = 3.14
	fmt.Println("f is set to ", f)

	whatWasSaid, another := saySomething("s1 string", "s2 string")
	fmt.Println("The function returns ", whatWasSaid, another)
}

func saySomething(s1 string, s2 string) (string, string) {
	return s1, s2
}
