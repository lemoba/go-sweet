package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// go doc net/http | grep "^func" | wc -l
// go doc net/http | grep "^type" | grep struct

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
