package git_cmd_worktree

import (
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
)

func (wc *WorktreeCommands) PruneWorktrees() error {
	_, err := os_commands.ExecOsCmd("git", "worktree", "prune")
	if err != nil {
		return err
	}
	return nil
}
