package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "battlesnake-go",
	Short: "Go Battlesnake, go!",
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
