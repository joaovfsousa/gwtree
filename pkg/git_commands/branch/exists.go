package git_cmd_branch

import "github.com/joaovfsousa/gwtree/internal/os_commands"

func BranchExists(branchName string) bool {
	_, err := os_commands.ExecOsCmd("git", "rev-parse", "--verify", branchName)

	return err == nil
}
