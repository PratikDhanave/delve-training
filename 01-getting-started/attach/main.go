package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})
	fmt.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
