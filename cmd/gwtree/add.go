package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/pkg/usecases"
)

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
		var err error

		newBranchName := args[0]

		baseBranchName := ""
		if len(args) > 1 {
			baseBranchName = args[1]
		} else {
			baseBranchName, err = uc.PickBranch()
			if err != nil {
				panic(err)
			}
		}

		err = uc.CreateWorktree(&usecases.CreateWorktreeOptions{
			NewBranchName:  newBranchName,
			BaseBranchName: baseBranchName,
		})
		if err != nil {
			l.Error(fmt.Sprintf("Failed to create worktree: %v", err.Error()))
		}
	},
}
