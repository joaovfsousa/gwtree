package usecases

import (
	"regexp"

	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
)

type CreateWorktreeOptions struct {
	NewBranchName  string
	BaseBranchName string
}

func getTreeNameFromBranchName(branchName string) string {
	return regexp.MustCompile("/").ReplaceAllString(branchName, "")
}

func (uc *UseCases) CreateWorktree(opts *CreateWorktreeOptions) error {
	treeName := getTreeNameFromBranchName(opts.NewBranchName)

	err := uc.gc.Worktree.AddWorktree(&git_cmd_worktree.WorktreeAddOptions{
		NewBranchName:  opts.NewBranchName,
		BaseBranchName: opts.BaseBranchName,
		TreeName:       treeName,
	})
	if err != nil {
		return err
	}

	return nil
}
