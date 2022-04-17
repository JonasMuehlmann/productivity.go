package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	tagCmd.AddCommand(shortCmd)

}
