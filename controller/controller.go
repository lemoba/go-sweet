package controller

import (
	"context"
	"fmt"
	"github.com/lemoba/go-sweet/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan any, 1)
	panicChan := make(chan any, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "middleware out")
		c.SetHasTimeout()
	}
	return nil
}

func UserLoginController(c *framework.Context) error {
	fmt.Println("user login")
	return c.Json(200, "success")
}

type User struct {
	Name  string `json:"name"`
	Age   uint8  `json:"age"`
	Email string `json:"email"`
}

func UserListController(c *framework.Context) error {
	fmt.Println("user list")

	list := []User{
		{
			"ranen", 12, "129@qq.com",
		},
		{
			"golang", 13, "golang@gmail.com",
		},
	}
	return c.Json(200, list)
}
