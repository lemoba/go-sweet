package controller

import (
	"fmt"
	"github.com/lemoba/go-sweet/provider/demo"
	"time"

	"github.com/lemoba/go-sweet/framework/gin"
)

func FooControllerHandler(c *gin.Context) {
	//finish := make(chan any, 1)
	//panicChan := make(chan any, 1)
	//
	//durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	//defer cancel()
	//
	//go func() {
	//	defer func() {
	//		if p := recover(); p != nil {
	//			panicChan <- p
	//		}
	//	}()
	//
	//	time.Sleep(10 * time.Second)
	//	c.ISetOkStatus().IJson("ok")
	//
	//	finish <- struct{}{}
	//}()
	//
	//select {
	//case p := <-panicChan:
	//	c.WriterMux().Lock()
	//	defer c.WriterMux().Unlock()
	//	log.Println(p)
	//	c.SetStatus(500).Json("panic")
	//case <-finish:
	//	fmt.Println("finish")
	//case <-durationCtx.Done():
	//	c.WriterMux().Lock()
	//	defer c.WriterMux().Unlock()
	//	c.SetStatus(500).Json("middleware out")
	//	c.SetHasTimeout()
	//}
	//return nil
	demoService := c.MustMake(demo.Key).(demo.Service)

	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("ok, UserLoginController: " + foo)
}

type User struct {
	Name  string `json:"name"`
	Age   uint8  `json:"age"`
	Email string `json:"email"`
}

func UserListController(c *gin.Context) {
	fmt.Println("user list")

	list := []User{
		{
			"ranen", 12, "129@qq.com",
		},
		{
			"golang", 13, "golang@gmail.com",
		},
	}
	c.ISetOkStatus().IJson(list)
}
