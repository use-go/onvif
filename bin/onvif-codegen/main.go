// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"
	"github.com/juju/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

var (
	// Logger is a zerolog logger, that can be safely used from any part of the application.
	// It gathers the format and the output. The application can replace the default Logger
	// for an alternative that meets its own output.
	Logger = zerolog.
		New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		With().Timestamp().
		Logger()
)

func main() {
	cmd := &cobra.Command{
		Use:   "codegen",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("Missing sub-command")
		},
	}

	sdk := &cobra.Command{
		Use:   "sdk",
		Short: "Generate the files of a SDK package",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
			defer cancel()
			return codegenSdk(ctx, args[0], args[1])
		},
	}

	dump := &cobra.Command{
		Use:   "cli",
		Short: "Generate the file of a CLI dump module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
			defer cancel()
			return codegenDump(ctx, args[0], args[1])
		},
	}

	cmd.AddCommand(sdk, dump)

	if err := cmd.Execute(); err != nil {
		Logger.Fatal().Err(err).Msg("Aborting")
	} else {
		Logger.Info().Msg("Exiting")
	}
}

func getwd() string {
	path, _ := os.Getwd()
	return path
}
