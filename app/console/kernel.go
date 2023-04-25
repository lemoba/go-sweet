package console

import (
	"github.com/lemoba/go-sweet/app/console/command/demo"
	"github.com/lemoba/go-sweet/framework"
	"github.com/lemoba/go-sweet/framework/cobra"
	"github.com/lemoba/go-sweet/framework/command"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "sweet",
		Short: "sweet 命令",
		Long:  "sweet 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		RunE: func(c *cobra.Command, args []string) error {
			c.InitDefaultHelpCmd()
			return c.Help()
		},

		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
