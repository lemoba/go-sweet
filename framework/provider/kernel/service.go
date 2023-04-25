package kernel

import (
	"github.com/lemoba/go-sweet/framework/gin"
	"net/http"
)

type SweetKernelService struct {
	engin *gin.Engine
}

func NewSweetService(params ...any) (any, error) {
	httpEngine := params[0].(*gin.Engine)
	return &SweetKernelService{engin: httpEngine}, nil
}

func (s *SweetKernelService) HttpEngine() http.Handler {
	return s.engin
}
