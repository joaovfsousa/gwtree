package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	logger "github.com/joaovfsousa/gwtree/internal"
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
)

var l = logger.GetLogger()

var rootCmd = &cobra.Command{
	Use:   "gwtree",
	Short: "Helper to use git worktrees",
}

var addCmd = &cobra.Command{
	Use:     "add [branchName] [<sourceBranchName>]",
	Aliases: []string{"a"},
	Short:   "Add a worktree",
	Long:    "To create a worktree interactive, don't provide a <sourceBranchName>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("branchName is required")
		}

		if len(args) > 2 {
			return errors.New("add only accepts two arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		branch_exists := git_cmd_branch.Exists(args[0])

		l.Info(fmt.Sprintf("Branch exists: %t", branch_exists))
	},
}

func RootExecute() error {
	rootCmd.AddCommand(addCmd)
	return rootCmd.Execute()
}
