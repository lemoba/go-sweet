package command

import "github.com/lemoba/go-sweet/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
}
