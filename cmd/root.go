package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

var (
	debug bool
	gcp   bool
)

var rootCmd = &cobra.Command{
	Use:   "battlesnake-go",
	Short: "Go Battlesnake, go!",
}

func init() {
	cobra.OnInitialize(initLogging)
	rootCmd.AddCommand(serveCmd)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Emit debug messages.")
	rootCmd.PersistentFlags().BoolVar(&gcp, "gcp", false, "Format log messages for Google Cloud Platform.")
}

func initLogging() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if gcp {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.LevelFieldName = "severity"
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
