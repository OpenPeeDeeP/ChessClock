package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/OpenPeeDeeP/chessclock/chessclock"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var version string

var connection *grpc.ClientConn
var client chessclock.ChessClockClient
var mainLogger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	rootCmd.AddCommand(startCmd, stopCmd, tagsCmd, sheetsCmd, scheduleCmd, tallyCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func startClient(log zerolog.Logger) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		connection, err = grpc.Dial("localhost:4242", grpc.WithInsecure())
		if err != nil {
			log.Error().Err(err).Msg("Could not connect to the daemon")
			return err
		}
		client = chessclock.NewChessClockClient(connection)
		return nil
	}
}

func stopClient(cmd *cobra.Command, args []string) error {
	return connection.Close()
}

func parseDate(date string) (int64, error) {
	dateString := strings.Split(date, "/")
	if len(dateString) != 3 {
		return 0, errors.New("Date not in the proper format (YYYY/MM/DD)")
	}
	dateInts := make([]int, 0, 3)
	for _, ds := range dateString {
		i, err := strconv.Atoi(ds)
		if err != nil {
			return 0, errors.New("Date not in the proper format (YYYY/MM/DD)")
		}
		dateInts = append(dateInts, i)
	}
	dateTime := time.Date(dateInts[0], time.Month(dateInts[1]), dateInts[2], 0, 0, 0, 0, time.Local).Unix()
	return dateTime, nil
}
