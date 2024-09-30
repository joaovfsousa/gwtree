package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gwtree",
	Short: "Helper to use git worktrees",
	Long:  "Helper to use git worktrees. Some commands require shell integration to be setup in order to make them work",
}

func RootExecute() error {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(switchCmd)
	return rootCmd.Execute()
}
