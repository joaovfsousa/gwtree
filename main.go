package main

import (
	"fmt"
	"log/slog"

	"github.com/joaovfsousa/gwtree/pkg/git_commands"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	worktrees, _ := git_commands.ListWorktrees(logger)

	for _, worktree := range worktrees {
		logger.Info(fmt.Sprintf("Path: %v  Branch: %v", worktree.Path, worktree.BranchName))
	}
}
