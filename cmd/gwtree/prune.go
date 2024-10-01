package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	git_cmd_worktree "github.com/joaovfsousa/gwtree/pkg/git_commands/worktree"
	"github.com/joaovfsousa/gwtree/pkg/usecases"
)

var dryRunFlagName = "dry-run"

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Aliases: []string{"p"},
	Short:   "Prunes old worktrees",
	Long:    "Prunes old worktrees. Shell integration has to be setup in order to make it work",
	Run: func(cmd *cobra.Command, _ []string) {
		FIFTEEN_DAYS_IN_SECONDS := int64(((24 * time.Hour) * 15).Seconds())

		dryRun, err := cmd.Flags().GetBool(dryRunFlagName)
		if err != nil {
			dryRun = false
		}

		wts, err := git_cmd_worktree.ListWorktrees()
		if err != nil {
			panic(err)
		}

		for _, wt := range wts {
			if wt.BranchName == "master" || wt.BranchName == "main" {
				continue
			}

			unixTime, err := git_cmd_worktree.GetLastWorktreeCommitDate(wt)
			if err != nil {
				panic(err)
			}

			t := time.Unix(unixTime, 0)

			if !(time.Now().Unix()-unixTime < FIFTEEN_DAYS_IN_SECONDS) {
				if dryRun {
					fmt.Printf("[%v] %v => Last commit on %v\n", wt.BranchName, wt.Path, t)
				} else {
					if err := usecases.DeleteWorktree(wt); err != nil {
						fmt.Printf("Failed to delete worktree [%v] %v", wt.BranchName, wt.Path)
					} else {
						fmt.Printf("Successfully deleted worktree [%v] %v", wt.BranchName, wt.Path)
					}
				}
			}
		}
	},
}
