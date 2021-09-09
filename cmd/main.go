package main

import (
	"github.com/spf13/cobra"
	"github.com/zag07/gin-example/cmd/serve"
)

// @title gin-example
// @version 0.3.x
func main() {
	rootCmd := newRootCmd()
	cobra.CheckErr(rootCmd.Execute())
}

func newRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "ginE",
		Short: "Gin example with Cobra",
		Long:  `A toy project`,
	}

	serve.RegisterCommandRecursive(cmd)

	return cmd
}
