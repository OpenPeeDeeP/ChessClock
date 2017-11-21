package main

import (
	"context"
	"errors"
	"time"

	"github.com/OpenPeeDeeP/chessclock/chessclock"
	"github.com/spf13/cobra"
)

var startLogger = mainLogger.With().Str("cmd", "tags").Logger()
var startCmd = &cobra.Command{
	Use:      "start [TAG]",
	Short:    "Start a specific task",
	Args:     validateStartCmdArgs,
	RunE:     startCmdRun,
	PreRunE:  startClient(startLogger),
	PostRunE: stopClient,
}

var startDescription string

func init() {
	startCmd.Flags().StringVarP(&startDescription, "description", "d", "", "Description for the task")
}

func startCmdRun(cmd *cobra.Command, args []string) error {
	_, err := client.Start(context.Background(), &chessclock.StartRequest{
		Timestamp:   time.Now().Unix(),
		Tag:         args[0],
		Description: startDescription,
	})
	if err != nil {
		startLogger.Error().Err(err).Msg("Could not start task")
		return err
	}
	return nil
}

func validateStartCmdArgs(cmd *cobra.Command, args []string) (err error) {
	switch {
	case len(args) > 1:
		err = errors.New("Too many arguments")
		startLogger.Error().Err(err).Msg("Incorrect Arguments")
		return err
	case len(args) < 1:
		err = errors.New("Must sepcify a tag for the task")
		startLogger.Error().Err(err).Msg("Incorrect Arguments")
		return err
	default:
		return nil
	}
}
