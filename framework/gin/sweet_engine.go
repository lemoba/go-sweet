package gin

import "github.com/lemoba/go-sweet/framework"

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}

func (engine *Engine) Binds(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBinds(key string) bool {
	return engine.container.IsBind(key)
}
