package demo

import "github.com/lemoba/go-sweet/framework"

type DemoProvider struct {
	framework.ServiceProvider

	container framework.Container
}

func (p *DemoProvider) Register(container framework.Container) framework.NewInstance {
	return NewService
}

func (p *DemoProvider) Boot(container framework.Container) error {
	p.container = container
	return nil
}

func (p *DemoProvider) IsDefer() bool {
	return false
}

func (p *DemoProvider) Params(container framework.Container) []any {
	return []any{p.container}
}

func (p *DemoProvider) Name() string {
	return DemoKey
}
