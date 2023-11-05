package main

import (
    "log"
    "sort"
)

type User struct {
    FirstName string
    LastName string
}

func main() {
    // map
    myMapString := make(map[string]string)

    myMapString["first-dog"] = "Samson"
    myMapString["second-dog"] = "Cassie"

    log.Println(myMapString)
    log.Println(myMapString["first-dog"])
    log.Println(myMapString["second-dog"])

    myMapStruct := make(map[string]User)

    me := User {
        FirstName: "Alice",
        LastName: "Blice",
    }

    myMapStruct["me"] = me

    log.Println(myMapStruct)
    log.Println(myMapStruct["me"].FirstName)
    log.Println(myMapStruct["me"].LastName)

    var myNewVar float64
    myNewVar = 11.23

    log.Println(myNewVar)

    // slices
    var myString string
    myString = "Fish"

    log.Println(myString)

    var mySlice []string
    mySlice = append(mySlice, "John")
    mySlice = append(mySlice, "Mary")

    log.Println(mySlice)

    var mySliceInt []int
    mySliceInt = append(mySliceInt, 1)
    mySliceInt = append(mySliceInt, 12)
    mySliceInt = append(mySliceInt, 345)
    mySliceInt = append(mySliceInt, 21)

    sort.Ints(mySliceInt)

    log.Println(mySliceInt)

    myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    log.Println(myNumbers)
    log.Println(myNumbers[0:2])
    log.Println(myNumbers[:3])
    log.Println(myNumbers[4:])
}
