package app

import (
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/contract"
)

// HadeAppProvider 提供App的具体实现方法
type SweetAppProvider struct {
	BaseFolder string
}

// Register 注册SweetApp方法
func (s *SweetAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewSweetApp
}

func (s *SweetAppProvider) Boot(container framework.Container) error {
	return nil
}

func (s *SweetAppProvider) IsDefer() bool {
	return false
}

func (s *SweetAppProvider) Params(container framework.Container) []any {
	return []any{container, s.BaseFolder}
}

func (s *SweetAppProvider) Name() string {
	return contract.AppKey
}
