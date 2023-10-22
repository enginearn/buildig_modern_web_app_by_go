package main

import (
	"fmt"
	"net/http"
)

// ポート番号を定義する
const portNumber = ":8080"

// ハンドラ関数を定義する
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}

// ハンドラ関数を定義する
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	fmt.Fprintf(w, "This is About Page\n")
	fmt.Fprintf(w, "Sum: %d\n", sum)
}

// ハンドラ関数を定義する
func addValues(x, y int) int {
	sum := x + y
	return sum
}

func DivideByZero(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is DivideByZero Page\n")
	f, err := divideByZero(10, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d divided by %d is %d\n", 10, 10, f)
}

func divideByZero(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	sum := x / y
	return sum, nil
}

func main() {
	// ハンドラ関数をルートパスにマップする
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", DivideByZero)

	// サーバを開始する
	fmt.Println("Server is starting at port", portNumber)
	http.ListenAndServe(":8080", nil)
}
