package chessclock

import (
	"errors"

	"github.com/OpenPeeDeeP/chessclock/chessclock"
)

type EventType int

const (
	ETStart EventType = 1 + iota
	ETStop
)

type Eventer interface{}

type Event struct {
	Type    EventType
	Details Eventer
}

func (e *Event) MustGetStartDetails() *StartEvent {
	start, ok := e.Details.(*StartEvent)
	if !ok {
		panic(errors.New("Event is not a start event"))
	}
	return start
}

func (e *Event) MustGetStopDetails() *StopEvent {
	stop, ok := e.Details.(*StopEvent)
	if !ok {
		panic(errors.New("Event is not a stop event"))
	}
	return stop
}

func (e *Event) IsStart() bool {
	_, ok := e.Details.(*StartEvent)
	return ok
}

func (e *Event) IsStop() bool {
	_, ok := e.Details.(*StopEvent)
	return ok
}

type StartEvent struct {
	StartTime   int64
	Tag         string
	Description string
}

type StopEvent struct {
	StopTime int64
	Reason   chessclock.StopRequest_Reason
}

type ChessClockStore interface {
	Start(timestamp int64, tag, description string) error
	Stop(reason chessclock.StopRequest_Reason) error
	TimeSheets() ([]int64, error)
	Events(date int64) ([]*Event, error)
}
