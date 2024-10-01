package git_cmd_worktree

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
)

type WorktreeAddOptions struct {
	TreeName       string
	NewBranchName  string
	BaseBranchName string
}

func AddWorktree(opts *WorktreeAddOptions) error {
	branchExists := git_cmd_branch.BranchExists(opts.NewBranchName)

	if branchExists {
		// TODO
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
