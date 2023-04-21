package route

import (
	"github.com/lemoba/go-sweet/controller"
	"github.com/lemoba/go-sweet/framework"
)

func RegisterRouter(core *framework.Core) {
	core.Get("foo", controller.FooControllerHandler)
}
