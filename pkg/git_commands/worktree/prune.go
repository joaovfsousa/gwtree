package git_cmd_worktree

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
)

func PruneWorktrees() error {
	_, err := os_commands.ExecOsCmd("git", "worktree", "prune")
	if err != nil {
		return err
	}
	return nil
}
