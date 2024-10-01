package git_cmd_worktree

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func (wc *WorktreeCommands) ListWorktrees() ([]*domain.Worktree, error) {
	lines, err := os_commands.ExecOsCmd("git", "worktree", "list")
	if err != nil {
		return nil, err
	}

	worktrees := []*domain.Worktree{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "(bare)") {
			continue
		}

		pattern := `^(.+?)\s+\w+\s+\[(.+?)\]$`

		re := regexp.MustCompile(pattern)

		matches := re.FindStringSubmatch(line)

		if len(matches) < 3 {
			wc.logger.Debug(fmt.Sprintf("Failed to parse: '%v'", line))

			continue
		}

		path := matches[1]
		branchName := matches[2]

		worktrees = append(worktrees, &domain.Worktree{BranchName: branchName, Path: path})
	}

	return worktrees, nil
}
