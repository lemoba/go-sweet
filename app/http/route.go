package http

import (
	"github.com/lemoba/go-sweet/app/http/module/demo"
	"github.com/lemoba/go-sweet/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
