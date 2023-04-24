package main

import (
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/middleware"
)

func main() {
	t1 := middleware.Test1()
	t3 := middleware.Test3()
	t2 := middleware.Test2()

	t1(&framework.Context{})
	t3(&framework.Context{})
	t2(&framework.Context{})
}
