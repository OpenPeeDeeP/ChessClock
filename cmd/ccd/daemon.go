package main

import (
	"context"

	"github.com/OpenPeeDeeP/chessclock/chessclock"
)

//ChessClockDameon is the daemon that handles tasks
type ChessClockDameon struct{}

//Start starts a task
func (ccd *ChessClockDameon) Start(context.Context, *chessclock.StartRequest) (*chessclock.StartResponse, error) {
	return nil, nil
}

//Stop will stop the previous task
func (ccd *ChessClockDameon) Stop(context.Context, *chessclock.StopRequest) (*chessclock.StopResponse, error) {
	return nil, nil
}

//Schedule list all tasks and when they were started
func (ccd *ChessClockDameon) Schedule(context.Context, *chessclock.ScheduleRequest) (*chessclock.ScheduleResponse, error) {
	return nil, nil
}

//Tally list all the tasks and how long they were worked on
func (ccd *ChessClockDameon) Tally(context.Context, *chessclock.TallyRequest) (*chessclock.TallyResponse, error) {
	return nil, nil
}

//ListTimeSheets list all the available time sheets
func (ccd *ChessClockDameon) ListTimeSheets(context.Context, *chessclock.ListTimeSheetsRequest) (*chessclock.ListTimeSheetsResponse, error) {
	return nil, nil
}

//ListTags list all the tags for a given time sheet
func (ccd *ChessClockDameon) ListTags(context.Context, *chessclock.ListTagsRequest) (*chessclock.ListTagsResponse, error) {
	return nil, nil
}

//Version returns the version of the daemon
func (ccd *ChessClockDameon) Version(context.Context, *chessclock.VersionRequest) (*chessclock.VersionResponse, error) {
	return &chessclock.VersionResponse{
		Version: version,
	}, nil
}
