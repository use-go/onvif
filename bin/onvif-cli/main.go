// Copyright (c) 2022-2022 Jean-Francois SMIGIELSKI

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
		Use:   "main",
		Short: "OnVif command Line Interface",
		Long:  "CLI Client for OnVif devices",
		//Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("Missing sub-command")
		},
	}

	discover := &cobra.Command{
		Use:   "discover",
		Short: "Discover the local cameras",
		Long:  "Discover the local cameras on all the local network interfaces using the built-in Web Service discovery mechanism",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
			defer cancel()
			return discover(ctx)
		},
	}

	dump := &cobra.Command{
		Use:   "dump",
		Short: "Dump the configuration of the given camera",
		Long:  "Dump the configuration of the given camera, playing all the possible OnVif calls to explicitly check which are supported",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
			defer cancel()
			return details(ctx, args[0])
		},
	}

	cmd.AddCommand(discover, dump)

	if err := cmd.Execute(); err != nil {
		Logger.Fatal().Err(err).Msg("Aborting")
	} else {
		Logger.Info().Msg("Exiting")
	}
}

func runAll(ctx context.Context, hooks ...func(ctx2 context.Context)) {
	for _, fun := range hooks {
		if ctx.Err() != nil {
			return
		}
		fun(ctx)
	}
}
