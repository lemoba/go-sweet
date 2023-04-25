package kernel

import (
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/contract"
	"github.com/lemoba/go-sweet/framework/gin"
)

type SweetKernelProvider struct {
	HttpEngine *gin.Engine
}

func (s *SweetKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewSweetService
}

func (s *SweetKernelProvider) Boot(container framework.Container) error {
	if s.HttpEngine == nil {
		s.HttpEngine = gin.Default()
	}

	s.HttpEngine.SetContainer(container)
	return nil
}

func (s *SweetKernelProvider) IsDefer() bool {
	return false
}

func (s *SweetKernelProvider) Params(container framework.Container) []any {
	return []any{s.HttpEngine}
}

func (s *SweetKernelProvider) Name() string {
	return contract.KernelKey
}
