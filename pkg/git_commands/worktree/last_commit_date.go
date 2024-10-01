package git_cmd_worktree

import (
	"fmt"
	"strconv"

	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func GetLastWorktreeCommitDate(wt *domain.Worktree) (int64, error) {
	lines, err := os_commands.ExecOsCmd("git", "-C", wt.Path, "log", "-1", "--format=%ct")
	if err != nil {
		return 0, err
	}

	if len(lines) < 1 {
		return 0, fmt.Errorf("Failed get last commit date of worktree [%v] at %v", wt.BranchName, wt.Path)
	}

	unixDateAsStr := lines[0]

	if len(unixDateAsStr) == 0 {
		return 0, fmt.Errorf("Failed to parse last commit date of worktree [%v] at %v", wt.BranchName, wt.Path)
	}

	unixDate, err := strconv.ParseInt(unixDateAsStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return unixDate, nil
}
