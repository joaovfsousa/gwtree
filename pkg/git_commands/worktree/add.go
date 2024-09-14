package git_cmd_worktree

import (
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
)

type WorktreeAddOptions struct {
	TreeName       string
	NewBranchName  string
	BaseBranchName string
}

func Add(opts *WorktreeAddOptions) error {
	branchExists := git_cmd_branch.Exists(opts.NewBranchName)

	if branchExists {
	}
	if opts.BaseBranchName == "" {
		_, err := os_commands.ExecOsCmd("git", "worktree", "add", opts.TreeName, opts.NewBranchName)
		if err != nil {
			return err
		}

		return nil
	}

	_, err := os_commands.ExecOsCmd("git", "worktree", "add", "-b", opts.NewBranchName, "--checkout", opts.TreeName, opts.BaseBranchName)
	if err != nil {
		return err
	}

	return nil
}
