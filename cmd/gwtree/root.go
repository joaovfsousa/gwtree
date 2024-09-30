package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gwtree",
	Short: "Helper to use git worktrees",
}

func RootExecute() error {
	rootCmd.AddCommand(addCmd)
	return rootCmd.Execute()
}
