package git_cmd_branch

import "log/slog"

type BranchCommands struct {
	Logger *slog.Logger
}

func CreateBranchCommands(logger *slog.Logger) *BranchCommands {
	return &BranchCommands{logger}
}
