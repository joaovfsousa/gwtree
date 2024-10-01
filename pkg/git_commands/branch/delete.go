package git_cmd_branch

import (
	"github.com/joaovfsousa/gwtree/internal/os_commands"
	"github.com/joaovfsousa/gwtree/pkg/domain"
)

func (*BranchCommands) DeleteBranch(branch *domain.Branch) error {
	_, err := os_commands.ExecOsCmd("git", "branch", "-q", "-D", branch.Name)
	if err != nil {
		return err
	}

	return nil
}
