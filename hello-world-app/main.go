package main

import (
	"fmt"
	"net/http"
	os "os"
)

func main() {
	// ハンドラ関数を定義する
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(os.Stdout, "Written %d bytes\n", []any{n}...)
	}

	// ハンドラ関数をルートパスにマップする
	http.HandleFunc("/", helloHandler)

	// サーバを開始する
	http.ListenAndServe(":8080", nil)
}
