package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ungood/battlesnake-go/server"
)

var port int
var hostname string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the battlesnake on the specified port.",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(hostname, port)
	},
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen on.")
	serveCmd.Flags().StringVarP(&hostname, "hostname", "n", "localhost", "Hostname to listen on.")
}
