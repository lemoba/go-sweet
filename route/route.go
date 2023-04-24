package route

import (
	"github.com/lemoba/go-sweet/controller"
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/middleware"
)

func RegisterRouter(core *framework.Core) {
	core.Get("foo", controller.FooControllerHandler)
	core.Get("/user/login", controller.UserLoginController)

	api := core.Group("/api")
	api.Use(middleware.Test2())
	{
		api.Get("/user/list", controller.UserListController)
	}
}
