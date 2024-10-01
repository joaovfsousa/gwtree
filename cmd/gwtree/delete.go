package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/pkg/domain"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [<branchName>]",
	Aliases: []string{"r", "d"},
	Short:   "Deletes a worktree",
	Long:    "To delete a worktree interactive, don't provide a <branchName>. Shell integration has to be setup in order to make it work",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && len(args[0]) == 0 {
			return errors.New("<branchName> can't be empty")
		}

		if len(args) > 1 {
			return errors.New("switch only accepts 1 argument: <branchName>")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var wt *domain.Worktree

		if len(args) == 1 {
			branchName := args[0]

			wt, err = uc.PickWorktree(&branchName)
		} else {
			wt, err = uc.PickWorktree(nil)
		}

		if err != nil {
			panic(err)
		}

		if err := uc.DeleteWorktree(wt); err != nil {
			panic(err)
		}
	},
}
