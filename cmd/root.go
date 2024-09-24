package cmd

import (
	"fmt"

	"github.com/mehulambastha/fschecker-cli/internal"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a directory for file changes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting watch for: ", args[0])
		internal.WatchDir(args[0])
	},
}

func Execute() {
	var rootCmd = &cobra.Command{Use: "fschecker"}
	rootCmd.AddCommand(watchCmd)
	rootCmd.Execute()
}
