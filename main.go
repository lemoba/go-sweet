package main

import (
	"github.com/lemoba/go-sweet/framework/middleware"
	"github.com/lemoba/go-sweet/route"
	"net/http"

	"github.com/lemoba/go-sweet/framework"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery())

	route.RegisterRouter(core)
	server := &http.Server{
		Addr:    ":8888",
		Handler: core,
	}
	server.ListenAndServe()
}
