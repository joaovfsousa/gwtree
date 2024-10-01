package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/internal/file_ops"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

var switchCmd = &cobra.Command{
	Use:     "switch [<branchName>]",
	Aliases: []string{"s"},
	Short:   "Switches to worktree",
	Long:    "To switch to a worktree interactive, don't provide a <branchName>. Shell integration has to be setup in order to make it work",
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
			fmt.Printf("Failed to pick a worktree: %v\n", err.Error())
			os.Exit(1)
		}

		if err := file_ops.RecordNewDir(wt.Path); err != nil {
			fmt.Printf("Failed to set new dir: %v\n", err.Error())
			os.Exit(1)
		}
	},
}
