package git_commands

import (
	"log/slog"

	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
)

type GitCommander struct {
	Branch   *git_cmd_branch.BranchCommands
	Worktree *git_cmd_worktree.WorktreeCommands
}

func CreateGitCommander(logger *slog.Logger) *GitCommander {
	bc := git_cmd_branch.CreateBranchCommands(logger)
	return &GitCommander{
		Branch:   bc,
		Worktree: git_cmd_worktree.CreateWorktreeCommands(logger, bc),
	}
}
