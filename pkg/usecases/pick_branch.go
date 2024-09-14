package usecases

import (
	git_cmd_branch "github.com/joaovfsousa/gwtree/pkg/git_commands/branch"
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
)

func PickBranch() (string, error) {
	branches, err := git_cmd_branch.List()
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
