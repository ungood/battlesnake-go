package cmd

import (
	stdlog "log"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

var (
	debug bool
	json  bool
)

var rootCmd = &cobra.Command{
	Use:   "battlesnake-go",
	Short: "Go Battlesnake, go!",
}

func init() {
	cobra.OnInitialize(initLogging)
	rootCmd.AddCommand(serveCmd)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Emit debug messages.")
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "Format log messages as JSON.")
}

func initLogging() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if json {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.LevelFieldName = "severity"
	} else {
		writer := zerolog.ConsoleWriter{
			Out:              os.Stdout,
			FormatFieldName:  func(i interface{}) string { return "" },
			FormatFieldValue: func(i interface{}) string { return "" },
		}

		log.Logger = log.Output(writer)
	}

	// Redirect standard library log to zerolog.
	stdlog.SetFlags(0)
	stdlog.SetOutput(log.Logger)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
