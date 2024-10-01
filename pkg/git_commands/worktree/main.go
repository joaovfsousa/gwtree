package git_cmd_worktree

import (
	"log/slog"

	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
)

type WorktreeCommands struct {
	logger *slog.Logger
	bc     *git_cmd_branch.BranchCommands
}

func CreateWorktreeCommands(logger *slog.Logger, bc *git_cmd_branch.BranchCommands) *WorktreeCommands {
	return &WorktreeCommands{logger, bc}
}
