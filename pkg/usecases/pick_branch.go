package usecases

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/git_commands"
)

func PickBranch(gc *git_commands.GitCommander) (string, error) {
	branches, err := gc.Branch.ListBranches()
	if err != nil {
		return "", nil
	}

	branch_names := []string{}
	for _, b := range branches {
		branch_names = append(branch_names, b.String())
	}

	selectedBranch, err := os_commands.FzfSelect(branch_names)
	if err != nil {
		return "", nil
	}

	return selectedBranch, nil
}
