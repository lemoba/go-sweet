package main

import (
	"github.com/lemoba/go-sweet/app/console"
	"github.com/lemoba/go-sweet/app/http"
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/provider/app"
	"github.com/lemoba/go-sweet/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewSweetContainer()

	// 绑定App服务提供者
	container.Bind(&app.SweetAppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.SweetKernelProvider{HttpEngine: engine})
	}
	// 运行root命令
	console.RunCommand(container)
}
