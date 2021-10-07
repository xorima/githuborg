package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "githuborg",
	Short: "A tool for manging your GitHub orginisation",
	Long: `githuborg allows you to easily manage your GitHub organisation.
	For example it can mass delete or approve branches/pull requests
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("org", "o", "", "Name of the GitHub org")
	rootCmd.MarkPersistentFlagRequired("org")
}
