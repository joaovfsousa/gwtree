package git_commands

import (
	"fmt"
	"log/slog"
	"os/exec"
	"regexp"
	"strings"

	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func ListWorktrees(logger *slog.Logger) ([]*domain.Worktree, error) {
	cmd := exec.Command("git", "worktree", "list")

	output, err := cmd.Output()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to execute command: %v", err))
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

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
			logger.Debug(fmt.Sprintf("Failed to parse: '%v'", line))

			continue
		}

		path := matches[1]
		branchName := matches[2]

		worktrees = append(worktrees, &domain.Worktree{BranchName: branchName, Path: path})
	}

	return worktrees, nil
}
