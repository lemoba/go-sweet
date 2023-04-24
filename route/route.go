package route

import (
	"github.com/lemoba/go-sweet/controller"
	"github.com/lemoba/go-sweet/framework/gin"
	"github.com/lemoba/go-sweet/framework/middleware"
)

func RegisterRouter(core *gin.Engine) {
	core.GET("foo", controller.FooControllerHandler)
	core.GET("/user/login", controller.UserLoginController)

	api := core.Group("/api")
	api.Use(middleware.Test2())
	{
		api.GET("/user/list", controller.UserListController)
	}
}
