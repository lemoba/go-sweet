package demo

import (
	"fmt"
	"github.com/lemoba/go-sweet/framework"
)

type DemoService struct {
	// 实现接口
	Service

	// 参数
	c framework.Container
}

func NewDemoService(params ...any) (any, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	return &DemoService{c: c}, nil
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}
