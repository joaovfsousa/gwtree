package git_cmd_branch

import (
	"github.com/joaovfsousa/gwtree/pkg/domain"
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
)

func List() ([]*domain.Branch, error) {
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
