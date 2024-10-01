package usecases

import (
	"os"

	"github.com/joaovfsousa/gwtree/pkg/domain"
	"github.com/joaovfsousa/gwtree/pkg/git_commands"
)

func DeleteWorktree(gc *git_commands.GitCommander, wt *domain.Worktree) error {
	if err := os.RemoveAll(wt.Path); err != nil {
		return err
	}

	if err := gc.Worktree.PruneWorktrees(); err != nil {
		return err
	}

	if err := gc.Branch.DeleteBranch(&domain.Branch{Name: wt.BranchName}); err != nil {
		return err
	}

	return nil
}
