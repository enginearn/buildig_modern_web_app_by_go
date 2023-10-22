package main

import (
	"channels/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(values chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	values <- randomNumber
}

func main() {
	// intChan := make(chan int, 3)
	// defer close(intChan)
	// intChan <- 1
	// intChan <- 2
	// intChan <- 3
	// close(intChan)

	// for v := range intChan {
	// 	println(v)
	// }

	values := make(chan int)
	defer close(values)
	go CalculateValue(values)
	value := <-values
	log.Println(value)
}

