package git_cmd_branch

import (
	"github.com/joaovfsousa/gwtree/pkg/os_commands"
)

func Exists(branchName string) bool {
	_, err := os_commands.ExecOsCmd("git", "rev-parse", "--verify", branchName)

	return err == nil
}
