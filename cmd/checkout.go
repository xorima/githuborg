package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "checkout all repos in the GitHub Orginisation",
	Long:  `Allows you to easily checkout all repositories in the GitHub Orginisation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("checkout called")
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
