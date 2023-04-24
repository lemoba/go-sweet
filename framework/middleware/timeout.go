package middleware

import (
	"context"
	"fmt"
	"github.com/lemoba/go-sweet/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panichan := make(chan any, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panichan <- p
				}

				// 使用next执行具体的业务逻辑
				c.Next()

				finish <- struct{}{}
			}()
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panichan:
			c.SetStatus(500).Json("time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json("time out")
		}

		return nil
	}
}
