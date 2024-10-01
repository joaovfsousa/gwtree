package usecases

import (
	"regexp"

	"github.com/joaovfsousa/gwtree/pkg/git_commands"
	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
)

type CreateWorktreeOptions struct {
	NewBranchName  string
	BaseBranchName string
}

func getTreeNameFromBranchName(branchName string) string {
	return regexp.MustCompile("/").ReplaceAllString(branchName, "")
}

func CreateWorktree(gc *git_commands.GitCommander, opts *CreateWorktreeOptions) error {
	treeName := getTreeNameFromBranchName(opts.NewBranchName)

	err := gc.Worktree.AddWorktree(&git_cmd_worktree.WorktreeAddOptions{
		NewBranchName:  opts.NewBranchName,
		BaseBranchName: opts.BaseBranchName,
		TreeName:       treeName,
	})
	if err != nil {
		return err
	}

	return nil
}
