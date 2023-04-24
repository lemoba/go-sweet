package demo

const Key = "sweet:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
