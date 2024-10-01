package usecases

import (
	"os"

	"github.com/joaovfsousa/gwtree/pkg/domain"
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
)

func DeleteWorktree(wt *domain.Worktree) error {
	if err := os.RemoveAll(wt.Path); err != nil {
		return err
	}

	if err := git_cmd_worktree.PruneWorktrees(); err != nil {
		return err
	}

	if err := git_cmd_branch.DeleteBranch(&domain.Branch{Name: wt.BranchName}); err != nil {
		return err
	}

	return nil
}
