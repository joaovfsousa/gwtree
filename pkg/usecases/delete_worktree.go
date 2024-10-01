package usecases

import (
	"os"

	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func (uc *UseCases) DeleteWorktree(wt *domain.Worktree) error {
	if err := os.RemoveAll(wt.Path); err != nil {
		return err
	}

	if err := uc.gc.Worktree.PruneWorktrees(); err != nil {
		return err
	}

	if err := uc.gc.Branch.DeleteBranch(&domain.Branch{Name: wt.BranchName}); err != nil {
		return err
	}

	return nil
}
