package main

import (
	"fmt"
	"net/http"
)

const port = ":8082"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user service /")
	})
	fmt.Printf("user service listening: %s", port)
	http.ListenAndServe(port, nil)
}
