package main

import (
	"fmt"
	"net/http"
)

const port = ":8083"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "cart service /")
	})
	fmt.Printf("cart service listening: %s", port)
	http.ListenAndServe(port, nil)
}
