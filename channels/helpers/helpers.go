package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(number int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	value := rand.Intn(number)
	return value
}
