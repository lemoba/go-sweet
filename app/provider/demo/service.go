package demo

import "github.com/lemoba/go-sweet/framework"

type Service struct {
	container framework.Container
}

func NewService(params ...any) (any, error) {
	container := params[0].(framework.Container)
	return &Service{container: container}, nil
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "foo",
		},
		{
			ID:   2,
			Name: "bar",
		},
	}
}
