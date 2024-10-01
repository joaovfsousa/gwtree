package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/pkg/domain"
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
	"github.com/joaovfsousa/gwtree/pkg/usecases"
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

			wt, err = usecases.PickWorktree(&branchName)
		} else {
			wt, err = usecases.PickWorktree(nil)
		}

		if err != nil {
			panic(err)
		}

		if err := os.RemoveAll(wt.Path); err != nil {
			panic(err)
		}

		if err := git_cmd_worktree.PruneWorktrees(); err != nil {
			panic(err)
		}

		if err := git_cmd_branch.DeleteBranch(&domain.Branch{Name: wt.BranchName}); err != nil {
			panic(err)
		}
	},
}
