package main

import (
	"log"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// switch
	var myInput string
	fmt.Println("Enter some word: ")
	fmt.Scanln(&myInput)
	myInput = strings.TrimSpace(myInput)  // 改行文字を削除

	switch myInput {
	case "cat":
		log.Println("myInput is set to", myInput)

	case "dog":
		log.Println("myInput is set to", myInput)

	default:
		log.Println("myInput is set to something else...")
	}

    // fmt.Scan, fmt.ScanLn: 短い入力や数値、単語などを受け取る場合に便利
	var input string
	fmt.Print("Please enter some text: ")
	fmt.Scanln(&input)
	fmt.Printf("You entered: %s\n", input)

	// bufio: 行単位での長い文字列やスペースを含む文字列を受け取る場合に便利
    var another_input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter some text: ")
	another_input, _ = reader.ReadString('\n')
	fmt.Printf("You entered: %s", another_input)
}
