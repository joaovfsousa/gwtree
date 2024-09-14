package usecases

import (
	"regexp"
	"strings"

	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
)

type CreateWorktreeOptions struct {
	NewBranchName  string
	BaseBranchName string
}

func getTreeNameFromBranchName(branchName string) string {
	parts := strings.Split(branchName, "/")

	partsLastIndex := len(parts) - 1
	partsLastValue := parts[partsLastIndex]

	parts[partsLastIndex] = regexp.MustCompile("[^0-9]").ReplaceAllString(partsLastValue, "")

	return strings.Join(parts, "")
}

func CreateWorktree(opts *CreateWorktreeOptions) error {
	treeName := getTreeNameFromBranchName(opts.NewBranchName)

	err := git_cmd_worktree.Add(&git_cmd_worktree.WorktreeAddOptions{
		NewBranchName:  opts.NewBranchName,
		BaseBranchName: opts.BaseBranchName,
		TreeName:       treeName,
	})
	if err != nil {
		return err
	}

	return nil
}
