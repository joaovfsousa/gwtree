package usecases

import (
	"fmt"
	"regexp"

	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func (uc *UseCases) PickWorktree(branchName *string) (*domain.Worktree, error) {
	worktrees, err := uc.gc.Worktree.ListWorktrees()
	if err != nil {
		return nil, err
	}

	if branchName != nil {
		for _, wt := range worktrees {
			if wt.BranchName == *branchName {
				return wt, nil
			}
		}
	}

	opts := []string{}
	for _, wt := range worktrees {
		opts = append(opts, fmt.Sprintf("%v [%v]", wt.Path, wt.BranchName))
	}

	selectedOpt, err := os_commands.FzfSelect(opts, branchName)
	if err != nil {
		return nil, err
	}

	pattern := `^.+\[(.+?)\]$`

	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(selectedOpt)

	if len(matches) < 2 {
		return nil, fmt.Errorf("Couldn't parsed selected option from list of worktrees")
	}

	selectedBranchName := matches[1]

	for _, wt := range worktrees {
		if wt.BranchName == selectedBranchName {
			return wt, nil
		}
	}

	return nil, fmt.Errorf("Couldn't find worktree from selected options from list of worktrees")
}
