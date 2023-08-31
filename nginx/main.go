package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "in test /")
	})
	http.HandleFunc("/r1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "in test /r1")
	})
	fmt.Println("running main in: 8088")
	http.ListenAndServe(":8088", nil)
}
