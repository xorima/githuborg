package cmd

import (
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use: "branch",
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.PersistentFlags().StringP("topic", "t", "", "Topic for the search to filter on")
	branchCmd.MarkPersistentFlagRequired("topic")
	branchCmd.PersistentFlags().StringP("branch-name", "n", "", "Name of the branch you wish to affect")
	branchCmd.MarkPersistentFlagRequired("branch-name")
}
