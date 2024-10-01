package git_cmd_branch

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func ListBranches() ([]*domain.Branch, error) {
	lines, err := os_commands.ExecOsCmd("git", "branch", "--format=%(refname:short)")
	if err != nil {
		return nil, err
	}

	branches := []*domain.Branch{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		branches = append(branches, &domain.Branch{Name: line})
	}

	return branches, nil
}
