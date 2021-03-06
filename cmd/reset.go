package cmd

import (
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database.",
	Run: func(cmd *cobra.Command, args []string) {
		fileDb.Truncate()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
