package main

import (
	"fmt"
	"net/http"
)

func add(a, b int) int {
	return a + b + 1
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello CI/CD")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
