// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"bufio"
	"context"
	"github.com/juju/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"strings"
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

type Method struct {
	Name string

	// Can the request being sent without parameter (it can be automated, then)
	NoArg bool
}

func getMethods(sourceFile string) []Method {
	out := make([]Method, 0)

	fin, err := os.Open(sourceFile)
	if err != nil {
		Logger.Fatal().Err(err).Str("wd", getwd()).Str("file", sourceFile).Msg("Failed to open the configuration file")
	}
	defer func() { _ = fin.Close() }()

	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tokens := strings.Split(line, " ")
		method := tokens[0]
		opts := ""
		if len(tokens) > 1 {
			opts = tokens[1]
		}
		if method == "" || strings.HasPrefix(method, "#") {
			continue
		}

		out = append(out, Method{Name: method, NoArg: opts == "noarg"})
	}

	return out
}