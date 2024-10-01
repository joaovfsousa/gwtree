package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
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

		fmt.Printf("Threshold set to %v day(s)\n", thresholdInDays)

		wts, err := gc.Worktree.ListWorktrees()
		if err != nil {
			fmt.Printf("Failed to list worktrees: %v\n", err.Error())
			os.Exit(1)
		}

		for _, wt := range wts {
			if wt.BranchName == "master" || wt.BranchName == "main" {
				continue
			}

			unixTime, err := gc.Worktree.GetLastWorktreeCommitDate(wt)
			if err != nil {
				fmt.Printf("Failed to get %v last commit date: %v\n", wt.BranchName, err.Error())
				continue
			}

			t := time.Unix(unixTime, 0)

			if !(time.Now().Unix()-unixTime < threshold) {
				if dryRun {
					fmt.Printf("[%v] %v => Last commit on %v\n", wt.BranchName, wt.Path, t)
				} else {
					if err := uc.DeleteWorktree(wt); err != nil {
						fmt.Printf("Failed to delete worktree [%v] %v: %v\n", wt.BranchName, wt.Path, err.Error())
					} else {
						fmt.Printf("Successfully deleted worktree [%v] %v\n", wt.BranchName, wt.Path)
					}
				}
			}
		}
	},
}
