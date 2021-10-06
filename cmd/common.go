package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func getStringFlag(flagName string, cmd *cobra.Command) string {
	name, err := cmd.Flags().GetString(flagName)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
