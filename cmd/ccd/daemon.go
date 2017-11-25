package main

import (
	"context"

	"google.golang.org/grpc"

	cc "github.com/OpenPeeDeeP/chessclock"
	"github.com/OpenPeeDeeP/chessclock/chessclock"
	"google.golang.org/grpc/codes"
)

//ChessClockDaemon is the daemon that handles tasks
type ChessClockDaemon struct {
	store cc.Storer
}

//NewDaemon creates a new daemon using the specified store
func NewDaemon(store cc.Storer) *ChessClockDaemon {
	return &ChessClockDaemon{
		store: store,
	}
}

//Start starts a task
func (ccd *ChessClockDaemon) Start(ctx context.Context, req *chessclock.StartRequest) (*chessclock.StartResponse, error) {
	if req.GetTimestamp() == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Must specify a timestamp")
	}
	if req.GetTag() == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Must specify a tag")
	}
	err := ccd.store.Start(req.GetTimestamp(), req.GetTag(), req.GetDescription())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &chessclock.StartResponse{}, nil
}

//Stop will stop the previous task
func (ccd *ChessClockDaemon) Stop(ctx context.Context, req *chessclock.StopRequest) (*chessclock.StopResponse, error) {
	err := ccd.store.Stop(req.GetReason())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &chessclock.StopResponse{}, nil
}

//Schedule list all tasks and when they were started
func (ccd *ChessClockDaemon) Schedule(ctx context.Context, req *chessclock.ScheduleRequest) (*chessclock.ScheduleResponse, error) {
	if req.GetDate() == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Must specify a date for the timesheet")
	}
	events, err := ccd.store.Events(req.GetDate())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	schedule := getSchedule(events)
	return &chessclock.ScheduleResponse{
		Tasks: schedule,
	}, nil
}

//Tally list all the tasks and how long they were worked on
func (ccd *ChessClockDaemon) Tally(ctx context.Context, req *chessclock.TallyRequest) (*chessclock.TallyResponse, error) {
	if req.GetDate() == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Must specify a date for the timesheet")
	}
	events, err := ccd.store.Events(req.GetDate())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	tally := getTally(events)
	return &chessclock.TallyResponse{
		Tasks: tally,
	}, nil
}

//ListTimeSheets list all the available time sheets
func (ccd *ChessClockDaemon) ListTimeSheets(ctx context.Context, req *chessclock.ListTimeSheetsRequest) (*chessclock.ListTimeSheetsResponse, error) {
	timeSheets, err := ccd.store.TimeSheets()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	return &chessclock.ListTimeSheetsResponse{
		Dates: timeSheets,
	}, nil
}

//ListTags list all the tags for a given time sheet
func (ccd *ChessClockDaemon) ListTags(ctx context.Context, req *chessclock.ListTagsRequest) (*chessclock.ListTagsResponse, error) {
	if req.GetDate() == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Must specify a date for the timesheet")
	}
	events, err := ccd.store.Events(req.GetDate())
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	tags := getTags(events)
	return &chessclock.ListTagsResponse{
		Tags: tags,
	}, nil
}

//Version returns the version of the daemon
func (ccd *ChessClockDaemon) Version(context.Context, *chessclock.VersionRequest) (*chessclock.VersionResponse, error) {
	return &chessclock.VersionResponse{
		Version: version,
	}, nil
}

func getSchedule(events []*cc.Event) []*chessclock.ScheduleResponse_Task {
	schedule := make([]*chessclock.ScheduleResponse_Task, 0, len(events))
	for _, event := range events {
		if event.IsStart() {
			e := event.MustGetStartDetails()
			schedule = append(schedule, &chessclock.ScheduleResponse_Task{
				Timestamp:   e.StartTime,
				Tag:         e.Tag,
				Description: e.Description,
			})
		}
		if event.IsStop() {
			e := event.MustGetStopDetails()
			schedule = append(schedule, &chessclock.ScheduleResponse_Task{
				Timestamp: e.StopTime,
				Tag:       e.Reason.String(),
			})
		}
	}
	return schedule
}

func getTally(events []*cc.Event) []*chessclock.TallyResponse_Task {
	tags := make(map[string][]int64, len(events))
	description := make(map[string]string, len(events))
	var prevTag string
	for _, event := range events {
		if event.IsStart() {
			e := event.MustGetStartDetails()
			if prevTag != "" {
				tags[prevTag] = append(tags[prevTag], e.StartTime)
			}
			tags[e.Tag] = append(tags[e.Tag], e.StartTime)
			description[e.Tag] = e.Description
			prevTag = e.Tag
		}
		if event.IsStop() {
			e := event.MustGetStopDetails()
			if prevTag != "" {
				tags[prevTag] = append(tags[prevTag], e.StopTime)
			}
			tags[e.Reason.String()] = append(tags[e.Reason.String()], e.StopTime)
			prevTag = e.Reason.String()
		}
	}
	tagSpans := make([]*chessclock.TallyResponse_Task, 0, len(tags))
	for tag, times := range tags {
		if len(times)%2 == 1 {
			times = times[:len(times)-1]
		}
		var prevTime int64
		var totalTime int64
		for i, time := range times {
			if i%2 == 1 {
				prevTime = time
			} else {
				totalTime += time - prevTime
			}
		}
		tagSpan := &chessclock.TallyResponse_Task{
			Timespan: totalTime,
			Tag:      tag,
		}
		if desc, ok := description[tag]; ok {
			tagSpan.Description = desc
		}
		tagSpans = append(tagSpans, tagSpan)
	}
	return tagSpans
}

func getTags(events []*cc.Event) []string {
	tags := make(map[string]struct{}, len(events))
	for _, event := range events {
		if event.IsStart() {
			e := event.MustGetStartDetails()
			if _, ok := tags[e.Tag]; !ok {
				tags[e.Tag] = struct{}{}
			}
		}
		if event.IsStop() {
			e := event.MustGetStopDetails()
			if _, ok := tags[e.Reason.String()]; !ok {
				tags[e.Reason.String()] = struct{}{}
			}
		}
	}
	t := make([]string, 0, len(tags))
	for tag := range tags {
		t = append(t, tag)
	}
	return t
}
