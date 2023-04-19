package main

import (
	"github.com/lemoba/go-sweet/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8888",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
