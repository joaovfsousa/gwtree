package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/pkg/file_ops"
	"github.com/joaovfsousa/gwtree/pkg/usecases"
)

var addSwitchCmd = &cobra.Command{
	Use:     "add-switch [branchName] [<sourceBranchName>]",
	Aliases: []string{"as"},
	Short:   "Add a worktree and switch to it",
	Long:    "To create a worktree interactive, don't provide a <sourceBranchName>. Shell integration has to be setup in order to make it work",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("branchName is required")
		}

		if len(args) > 2 {
			return errors.New("add-switch only accepts two arguments")
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
				fmt.Printf("Failed to pick a base branch: %v\n", err.Error())
				os.Exit(1)
			}
		}

		err = uc.CreateWorktree(&usecases.CreateWorktreeOptions{
			NewBranchName:  newBranchName,
			BaseBranchName: baseBranchName,
		})
		if err != nil {
			fmt.Printf("Failed to create worktree: %v\n", err.Error())
			os.Exit(1)
		}

		wt, err := uc.PickWorktree(&newBranchName, false)
		if err != nil {
			fmt.Printf("Failed to find new worktree: %v\n", err.Error())
			os.Exit(1)
		}

		if err := file_ops.RecordNewDir(wt.Path); err != nil {
			fmt.Printf("Failed to set new dir: %v\n", err.Error())
			os.Exit(1)
		}
	},
}
