package cmd

import (
	"fmt"

	_ "github.com/mehulambastha/fschecker-cli/internal"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a directory for file changes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WATCHING YAY!")
	},
}

func Execute() {
	var rootCmd = &cobra.Command{Use: "fschecker"}
	rootCmd.AddCommand(watchCmd)
	rootCmd.Execute()
}
