package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/OpenPeeDeeP/chessclock/chessclock"
	"github.com/spf13/cobra"
)

var tallyLogger = mainLogger.With().Str("cmd", "tally").Logger()
var tallyCmd = &cobra.Command{
	Use:      "tally [DATE]",
	Short:    "Show a tally of how long a task was worked on for the day",
	Args:     validateTallyCmdArgs,
	RunE:     tallyCmdRun,
	PreRunE:  startClient(tallyLogger),
	PostRunE: stopClient,
}

func tallyCmdRun(cmd *cobra.Command, args []string) error {
	var date int64
	var err error
	if len(args) < 1 {
		date = time.Now().Unix()
	} else {
		date, err = parseDate(args[0])
		if err != nil {
			tallyLogger.Error().Err(err).Msg("How did you get here")
			return err
		}
	}
	res, err := client.Tally(context.Background(), &chessclock.TallyRequest{
		Date: date,
	})
	if err != nil {
		tallyLogger.Error().Err(err).Msg("Couldn't get a tally of all the tasks")
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprint(w, "DURATION\tTAG\tDESCRIPTION")
	for _, tally := range res.GetTasks() {
		ts := time.Duration(tally.GetTimespan())
		fmt.Fprintf(w, "%v\t%s\t%s", ts, tally.GetTag(), tally.GetDescription())
	}
	return nil
}

func validateTallyCmdArgs(cmd *cobra.Command, args []string) (err error) {
	if len(args) > 1 {
		err = errors.New("Too many arguments")
		tallyLogger.Error().Err(err).Msg("Incorrect Arguments")
		return err
	}
	if len(args) != 1 {
		return nil
	}
	if _, err = parseDate(args[0]); err != nil {
		tallyLogger.Error().Err(err).Msg("Incorrect Arguments")
		return err
	}
	return nil
}
