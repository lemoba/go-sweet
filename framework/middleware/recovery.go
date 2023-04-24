package middleware

import "github.com/lemoba/go-sweet/framework"

// recovery机制，将协程中的函数异常进行捕获
func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		defer func() {
			if err := Recovery(); err != nil {
				c.Json(500, err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()
		return nil
	}
}
