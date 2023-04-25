package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error

	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务
	Make(key string) (any, error)

	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) any

	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []any) (any, error)
}

type SweetContainer struct {
	Container

	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider

	// instance 存储具体的实例，key为字符串凭证
	instances map[string]any

	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// 创建一个服务容器
func NewSweetContainer() *SweetContainer {
	return &SweetContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]any{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (s *SweetContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range s.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}

	return ret
}

// Bind 将服务容器和关键字做了绑定
func (s *SweetContainer) Bind(provider ServiceProvider) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	key := provider.Name()

	s.providers[key] = provider

	if provider.IsDefer() == false {
		if err := provider.Boot(s); err != nil {
			return err
		}

		// 实例化方法
		params := provider.Params(s)
		method := provider.Register(s)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		s.instances[key] = instance
	}

	return nil
}

func (s *SweetContainer) findServiceProvider(key string) ServiceProvider {
	s.lock.RLock()
	defer s.lock.RLock()

	if sp, ok := s.providers[key]; ok {
		return sp
	}

	return nil
}

func (s *SweetContainer) IsBind(key string) bool {
	return s.findServiceProvider(key) != nil
}

func (s *SweetContainer) Make(key string) (any, error) {
	return s.make(key, nil, false)
}

func (s *SweetContainer) MustMake(key string) any {
	serv, err := s.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (s *SweetContainer) newInstance(sp ServiceProvider, params []any) (any, error) {
	if err := sp.Boot(s); err != nil {
		return nil, err
	}

	if params == nil {
		params = sp.Params(s)
	}

	method := sp.Register(s)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return ins, err
}

func (s *SweetContainer) make(key string, params []any, forceNew bool) (any, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := s.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract" + key + " have not register")
	}

	if forceNew {
		return s.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := s.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := s.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	s.instances[key] = inst
	return inst, nil
}
