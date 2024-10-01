package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/pkg/usecases"
)

const (
	dryRunFlagName    = "dry-run"
	thresholdFlagName = "threshold"
)

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Aliases: []string{"p"},
	Short:   "Prunes old worktrees",
	Long:    "Prunes old worktrees. Shell integration has to be setup in order to make it work",
	Run: func(cmd *cobra.Command, _ []string) {
		dryRun, err := cmd.Flags().GetBool(dryRunFlagName)
		if err != nil {
			dryRun = false
		}

		thresholdInDays, err := cmd.Flags().GetUint8(thresholdFlagName)
		if err != nil {
			thresholdInDays = 15
		}

		threshold := int64(((24 * time.Hour) * time.Duration(thresholdInDays)).Seconds())

		wts, err := gc.Worktree.ListWorktrees()
		if err != nil {
			panic(err)
		}

		for _, wt := range wts {
			if wt.BranchName == "master" || wt.BranchName == "main" {
				continue
			}

			unixTime, err := gc.Worktree.GetLastWorktreeCommitDate(wt)
			if err != nil {
				panic(err)
			}

			t := time.Unix(unixTime, 0)

			if !(time.Now().Unix()-unixTime < threshold) {
				if dryRun {
					fmt.Printf("[%v] %v => Last commit on %v\n", wt.BranchName, wt.Path, t)
				} else {
					if err := usecases.DeleteWorktree(gc, wt); err != nil {
						fmt.Printf("Failed to delete worktree [%v] %v", wt.BranchName, wt.Path)
					} else {
						fmt.Printf("Successfully deleted worktree [%v] %v", wt.BranchName, wt.Path)
					}
				}
			}
		}
	},
}
