package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	logger "github.com/joaovfsousa/gwtree/internal"
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
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
		branches, err := git_cmd_branch.List()
		if err != nil {
			panic(err)
		}

		branch_names := []string{}
		for _, b := range branches {
			branch_names = append(branch_names, b.String())
		}

		selectedBranch, err := os_commands.FzfSelect(branch_names)
		if err != nil {
			panic(err)
		}

		l.Info(fmt.Sprintf("Chosen: %v", selectedBranch))
	},
}

func RootExecute() error {
	rootCmd.AddCommand(addCmd)
	return rootCmd.Execute()
}
