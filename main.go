package main

import (
	"fmt"
	"log/slog"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	cmd := exec.Command("git", "worktree", "list")

	output, err := cmd.Output()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to execute command: %v", err))
	}

	lines := strings.Split(string(output), "\n")

	worktrees := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, "(bare)") {
			continue
		}

		worktrees = append(worktrees, line)
	}

	for _, line := range worktrees {
		pattern := `^(.+?)\s+\w+\s+\[(.+?)\]$`

		re := regexp.MustCompile(pattern)

		matches := re.FindStringSubmatch(line)

		if len(matches) < 3 {
			logger.Error(line)
			panic("Failed to list worktrees")
		}

		path := matches[1]
		branchName := matches[2]

		logger.Info(fmt.Sprintf("Path: %v; Branch: %v", path, branchName))
	}
}
