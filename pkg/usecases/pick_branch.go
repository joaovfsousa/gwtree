package usecases

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
)

func (uc *UseCases) PickBranch() (string, error) {
	branches, err := uc.gc.Branch.ListBranches()
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
