package cmd

import (
	"github.com/spf13/cobra"

	"github.com/joaovfsousa/gwtree/internal/logger"
	"github.com/joaovfsousa/gwtree/pkg/git_commands"
	"github.com/joaovfsousa/gwtree/pkg/usecases"
)

var rootCmd = &cobra.Command{
	Use:   "gwtree",
	Short: "Helper to use git worktrees",
	Long:  "Helper to use git worktrees. Some commands require shell integration to be setup in order to make them work",
}

var (
	l  = logger.GetLogger()
	gc = git_commands.CreateGitCommander(l)
	uc = usecases.CreateUseCases(l, gc)
)

const (
	noQueryFlagName        = "no-query"
	noQueryFlagDescription = "Sets if <branchName/worktreeName> should be used as fzf qeury in case a worktree or branch can't be found. Default = false"
)

func RootExecute() error {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.PersistentFlags().Bool(dryRunFlagName, false, "Log worktrees that will be deleted instead of actually deleting. Default = false")
	pruneCmd.PersistentFlags().Uint8(thresholdFlagName, 15, "Threshold in days. Default = 15")

	switchCmd.PersistentFlags().Bool(noQueryFlagName, false, noQueryFlagDescription)
	deleteCmd.PersistentFlags().Bool(noQueryFlagName, false, noQueryFlagDescription)

	return rootCmd.Execute()
}
