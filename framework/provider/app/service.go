package app

import (
	"errors"
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/util"
	flag "github.com/spf13/pflag"
	"path/filepath"
)

type SweetApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

func (s SweetApp) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (s SweetApp) BaseFolder() string {
	if s.baseFolder != "" {
		return s.baseFolder
	}

	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (s SweetApp) ConfigFolder() string {
	return filepath.Join(s.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (s SweetApp) LogFolder() string {
	return filepath.Join(s.BaseFolder(), "config")
}

func (s SweetApp) HttpFolder() string {
	return filepath.Join(s.BaseFolder(), "http")
}

func (s SweetApp) ConsoleFolder() string {
	return filepath.Join(s.BaseFolder(), "console")
}

func (s SweetApp) StorageFolder() string {
	return filepath.Join(s.BaseFolder(), "storage")
}

func (s SweetApp) ProviderFolder() string {
	return filepath.Join(s.BaseFolder(), "provider")
}

func (s SweetApp) MiddlewareFolder() string {
	return filepath.Join(s.HttpFolder(), "middleware")
}

func (s SweetApp) CommandFolder() string {
	return filepath.Join(s.ConsoleFolder(), "command")
}

func (s SweetApp) RuntimeFolder() string {
	return filepath.Join(s.StorageFolder(), "runtime")
}

func (s SweetApp) TestFolder() string {
	return filepath.Join(s.BaseFolder(), "test")
}

func NewSweetApp(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &SweetApp{
		baseFolder: baseFolder,
		container:  container,
	}, nil
}
