package git_cmd_worktree

import (
	"fmt"

	"github.com/joaovfsousa/gwtree/internal/os_commands"
)

type WorktreeAddOptions struct {
	TreeName       string
	NewBranchName  string
	BaseBranchName string
}

func (wc *WorktreeCommands) AddWorktree(opts *WorktreeAddOptions) error {
	branchExists := wc.bc.BranchExists(opts.NewBranchName)

	if branchExists {
		_, err := os_commands.ExecOsCmd("git", "worktree", "add", opts.TreeName, opts.NewBranchName)
		if err != nil {
			return err
		}

		return nil
	}

	baseBranchExists := wc.bc.BranchExists(opts.BaseBranchName)

	if !baseBranchExists {
		return fmt.Errorf("Base Branch %v doesn't exists", opts.BaseBranchName)
	}

	_, err := os_commands.ExecOsCmd("git", "worktree", "add", "-b", opts.NewBranchName, "--checkout", opts.TreeName, opts.BaseBranchName)
	if err != nil {
		return err
	}

	return nil
}
