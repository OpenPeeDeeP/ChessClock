// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chessclock.proto

/*
Package chessclock is a generated protocol buffer package.

It is generated from these files:
	chessclock.proto

It has these top-level messages:
	StartRequest
	StartResponse
	StopRequest
	StopResponse
	ScheduleRequest
	ScheduleResponse
	TallyRequest
	TallyResponse
	ListTimeSheetsRequest
	ListTimeSheetsResponse
	ListTagsRequest
	ListTagsResponse
*/
package chessclock

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StopRequest_Reason int32

const (
	StopRequest_Break    StopRequest_Reason = 0
	StopRequest_Lunch    StopRequest_Reason = 1
	StopRequest_EndOfDay StopRequest_Reason = 2
)

var StopRequest_Reason_name = map[int32]string{
	0: "Break",
	1: "Lunch",
	2: "EndOfDay",
}
var StopRequest_Reason_value = map[string]int32{
	"Break":    0,
	"Lunch":    1,
	"EndOfDay": 2,
}

func (x StopRequest_Reason) String() string {
	return proto.EnumName(StopRequest_Reason_name, int32(x))
}
func (StopRequest_Reason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type StartRequest struct {
	Timestamp   int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *StartRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *StartRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Empty response but useful if we need a resposne later.
type StartResponse struct {
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StopRequest struct {
	Reason StopRequest_Reason `protobuf:"varint,1,opt,name=reason,enum=chessclock.StopRequest_Reason" json:"reason,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StopRequest) GetReason() StopRequest_Reason {
	if m != nil {
		return m.Reason
	}
	return StopRequest_Break
}

// Empty response but useful if we need a resposne later.
type StopResponse struct {
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ScheduleRequest struct {
	Date int64 `protobuf:"varint,1,opt,name=date" json:"date,omitempty"`
}

func (m *ScheduleRequest) Reset()                    { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()               {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ScheduleRequest) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type ScheduleResponse struct {
	Tasks []*ScheduleResponse_Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *ScheduleResponse) Reset()                    { *m = ScheduleResponse{} }
func (m *ScheduleResponse) String() string            { return proto.CompactTextString(m) }
func (*ScheduleResponse) ProtoMessage()               {}
func (*ScheduleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ScheduleResponse) GetTasks() []*ScheduleResponse_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ScheduleResponse_Task struct {
	Timestamp   int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ScheduleResponse_Task) Reset()                    { *m = ScheduleResponse_Task{} }
func (m *ScheduleResponse_Task) String() string            { return proto.CompactTextString(m) }
func (*ScheduleResponse_Task) ProtoMessage()               {}
func (*ScheduleResponse_Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ScheduleResponse_Task) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ScheduleResponse_Task) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ScheduleResponse_Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type TallyRequest struct {
	Date int64 `protobuf:"varint,1,opt,name=date" json:"date,omitempty"`
}

func (m *TallyRequest) Reset()                    { *m = TallyRequest{} }
func (m *TallyRequest) String() string            { return proto.CompactTextString(m) }
func (*TallyRequest) ProtoMessage()               {}
func (*TallyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TallyRequest) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type TallyResponse struct {
	Tasks []*TallyResponse_Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TallyResponse) Reset()                    { *m = TallyResponse{} }
func (m *TallyResponse) String() string            { return proto.CompactTextString(m) }
func (*TallyResponse) ProtoMessage()               {}
func (*TallyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TallyResponse) GetTasks() []*TallyResponse_Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TallyResponse_Task struct {
	Timespan    int64  `protobuf:"varint,1,opt,name=timespan" json:"timespan,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *TallyResponse_Task) Reset()                    { *m = TallyResponse_Task{} }
func (m *TallyResponse_Task) String() string            { return proto.CompactTextString(m) }
func (*TallyResponse_Task) ProtoMessage()               {}
func (*TallyResponse_Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *TallyResponse_Task) GetTimespan() int64 {
	if m != nil {
		return m.Timespan
	}
	return 0
}

func (m *TallyResponse_Task) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *TallyResponse_Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Empty request but useful if we need parameters later.
type ListTimeSheetsRequest struct {
}

func (m *ListTimeSheetsRequest) Reset()                    { *m = ListTimeSheetsRequest{} }
func (m *ListTimeSheetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTimeSheetsRequest) ProtoMessage()               {}
func (*ListTimeSheetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListTimeSheetsResponse struct {
	Dates []int64 `protobuf:"varint,1,rep,packed,name=dates" json:"dates,omitempty"`
}

func (m *ListTimeSheetsResponse) Reset()                    { *m = ListTimeSheetsResponse{} }
func (m *ListTimeSheetsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTimeSheetsResponse) ProtoMessage()               {}
func (*ListTimeSheetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListTimeSheetsResponse) GetDates() []int64 {
	if m != nil {
		return m.Dates
	}
	return nil
}

type ListTagsRequest struct {
	Date int64 `protobuf:"varint,1,opt,name=date" json:"date,omitempty"`
}

func (m *ListTagsRequest) Reset()                    { *m = ListTagsRequest{} }
func (m *ListTagsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTagsRequest) ProtoMessage()               {}
func (*ListTagsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListTagsRequest) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type ListTagsResponse struct {
	Tags []string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
}

func (m *ListTagsResponse) Reset()                    { *m = ListTagsResponse{} }
func (m *ListTagsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTagsResponse) ProtoMessage()               {}
func (*ListTagsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListTagsResponse) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "chessclock.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "chessclock.StartResponse")
	proto.RegisterType((*StopRequest)(nil), "chessclock.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "chessclock.StopResponse")
	proto.RegisterType((*ScheduleRequest)(nil), "chessclock.ScheduleRequest")
	proto.RegisterType((*ScheduleResponse)(nil), "chessclock.ScheduleResponse")
	proto.RegisterType((*ScheduleResponse_Task)(nil), "chessclock.ScheduleResponse.Task")
	proto.RegisterType((*TallyRequest)(nil), "chessclock.TallyRequest")
	proto.RegisterType((*TallyResponse)(nil), "chessclock.TallyResponse")
	proto.RegisterType((*TallyResponse_Task)(nil), "chessclock.TallyResponse.Task")
	proto.RegisterType((*ListTimeSheetsRequest)(nil), "chessclock.ListTimeSheetsRequest")
	proto.RegisterType((*ListTimeSheetsResponse)(nil), "chessclock.ListTimeSheetsResponse")
	proto.RegisterType((*ListTagsRequest)(nil), "chessclock.ListTagsRequest")
	proto.RegisterType((*ListTagsResponse)(nil), "chessclock.ListTagsResponse")
	proto.RegisterEnum("chessclock.StopRequest_Reason", StopRequest_Reason_name, StopRequest_Reason_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChessClock service

type ChessClockClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	Tally(ctx context.Context, in *TallyRequest, opts ...grpc.CallOption) (*TallyResponse, error)
	ListTimeSheets(ctx context.Context, in *ListTimeSheetsRequest, opts ...grpc.CallOption) (*ListTimeSheetsResponse, error)
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
}

type chessClockClient struct {
	cc *grpc.ClientConn
}

func NewChessClockClient(cc *grpc.ClientConn) ChessClockClient {
	return &chessClockClient{cc}
}

func (c *chessClockClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessClockClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessClockClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/Schedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessClockClient) Tally(ctx context.Context, in *TallyRequest, opts ...grpc.CallOption) (*TallyResponse, error) {
	out := new(TallyResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/Tally", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessClockClient) ListTimeSheets(ctx context.Context, in *ListTimeSheetsRequest, opts ...grpc.CallOption) (*ListTimeSheetsResponse, error) {
	out := new(ListTimeSheetsResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/ListTimeSheets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessClockClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := grpc.Invoke(ctx, "/chessclock.ChessClock/ListTags", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChessClock service

type ChessClockServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	Tally(context.Context, *TallyRequest) (*TallyResponse, error)
	ListTimeSheets(context.Context, *ListTimeSheetsRequest) (*ListTimeSheetsResponse, error)
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
}

func RegisterChessClockServer(s *grpc.Server, srv ChessClockServer) {
	s.RegisterService(&_ChessClock_serviceDesc, srv)
}

func _ChessClock_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessClock_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessClock_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessClock_Tally_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).Tally(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/Tally",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).Tally(ctx, req.(*TallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessClock_ListTimeSheets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimeSheetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).ListTimeSheets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/ListTimeSheets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).ListTimeSheets(ctx, req.(*ListTimeSheetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessClock_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessClockServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chessclock.ChessClock/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessClockServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChessClock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chessclock.ChessClock",
	HandlerType: (*ChessClockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ChessClock_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ChessClock_Stop_Handler,
		},
		{
			MethodName: "Schedule",
			Handler:    _ChessClock_Schedule_Handler,
		},
		{
			MethodName: "Tally",
			Handler:    _ChessClock_Tally_Handler,
		},
		{
			MethodName: "ListTimeSheets",
			Handler:    _ChessClock_ListTimeSheets_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _ChessClock_ListTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chessclock.proto",
}

func init() { proto.RegisterFile("chessclock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x93, 0x4d, 0x5b, 0xb5, 0xd3, 0x6e, 0x37, 0x1a, 0x01, 0x1b, 0xc2, 0x0a, 0x15, 0x4b,
	0xa0, 0x1e, 0x50, 0x0e, 0x05, 0xc1, 0x01, 0x89, 0x03, 0x0b, 0xe2, 0xb2, 0x12, 0x92, 0x5b, 0x21,
	0xb8, 0x61, 0x52, 0xd3, 0x46, 0x6d, 0x93, 0x10, 0xbb, 0x87, 0x7d, 0x18, 0x8e, 0xbc, 0x02, 0xcf,
	0x87, 0xec, 0xb8, 0x4d, 0x1c, 0x4a, 0x0e, 0x48, 0xdc, 0xfc, 0xe7, 0xcb, 0xcc, 0xfc, 0x66, 0x3e,
	0x07, 0xfc, 0x78, 0xcd, 0x85, 0x88, 0xb7, 0x59, 0xbc, 0x89, 0xf2, 0x22, 0x93, 0x19, 0x42, 0x75,
	0x42, 0xbe, 0xc0, 0x68, 0x2e, 0x59, 0x21, 0x29, 0xff, 0xbe, 0xe7, 0x42, 0xe2, 0x15, 0x0c, 0x64,
	0xb2, 0xe3, 0x42, 0xb2, 0x5d, 0x1e, 0xb8, 0x13, 0x77, 0xea, 0xd1, 0xea, 0x00, 0x7d, 0xf0, 0x24,
	0x5b, 0x05, 0x67, 0x13, 0x77, 0x3a, 0xa0, 0x6a, 0x89, 0x13, 0x18, 0x2e, 0xb9, 0x88, 0x8b, 0x24,
	0x97, 0x49, 0x96, 0x06, 0x9e, 0xbe, 0xa9, 0x1f, 0x91, 0x0b, 0x38, 0x37, 0x19, 0x44, 0x9e, 0xa5,
	0x82, 0x13, 0x01, 0xc3, 0xb9, 0xcc, 0xf2, 0x43, 0xc6, 0x17, 0xd0, 0x2b, 0x38, 0x13, 0x59, 0xaa,
	0xd3, 0x8d, 0x67, 0x0f, 0xa3, 0x5a, 0xc1, 0x35, 0x61, 0x44, 0xb5, 0x8a, 0x1a, 0x35, 0x79, 0x0a,
	0xbd, 0xf2, 0x04, 0x07, 0xd0, 0x7d, 0x53, 0x70, 0xb6, 0xf1, 0x1d, 0xb5, 0xbc, 0xd9, 0xa7, 0xf1,
	0xda, 0x77, 0x71, 0x04, 0xfd, 0x77, 0xe9, 0xf2, 0xc3, 0xb7, 0xb7, 0xec, 0xd6, 0x3f, 0x23, 0x63,
	0xc5, 0xa9, 0x62, 0x99, 0x22, 0x1e, 0xc3, 0xc5, 0x3c, 0x5e, 0xf3, 0xe5, 0x7e, 0xcb, 0x0f, 0x85,
	0x20, 0x74, 0x96, 0x4c, 0x72, 0x43, 0xad, 0xd7, 0xe4, 0xa7, 0x0b, 0x7e, 0xa5, 0x2b, 0xbf, 0xc5,
	0x97, 0xd0, 0x95, 0x4c, 0x6c, 0x44, 0xe0, 0x4e, 0xbc, 0xe9, 0x70, 0xf6, 0xc8, 0x2a, 0xb8, 0x21,
	0x8e, 0x16, 0x4c, 0x6c, 0x68, 0xa9, 0x0f, 0x3f, 0x41, 0x47, 0x6d, 0xff, 0x43, 0x93, 0x09, 0x8c,
	0x16, 0x6c, 0xbb, 0xbd, 0x6d, 0x63, 0xf9, 0xe1, 0xc2, 0xb9, 0x11, 0x19, 0x90, 0xe7, 0x36, 0x88,
	0xd5, 0x79, 0x4b, 0x69, 0x51, 0x7c, 0x34, 0x14, 0x21, 0xf4, 0x75, 0xd1, 0x39, 0x4b, 0x4d, 0x9e,
	0xe3, 0xfe, 0x9f, 0x18, 0x2e, 0xe1, 0xee, 0x4d, 0x22, 0xe4, 0x22, 0xd9, 0xf1, 0xf9, 0x9a, 0x73,
	0x29, 0x0c, 0x0c, 0x89, 0xe0, 0x5e, 0xf3, 0xc2, 0x00, 0xdc, 0x81, 0xae, 0x42, 0x2b, 0x01, 0x3c,
	0x5a, 0x6e, 0xd4, 0x6c, 0xb5, 0x9e, 0xad, 0x44, 0x5b, 0x3f, 0x9e, 0x80, 0x5f, 0xc9, 0x4c, 0x40,
	0x84, 0x8e, 0x64, 0xab, 0x32, 0xde, 0x80, 0xea, 0xf5, 0xec, 0x97, 0x07, 0x70, 0xad, 0x1a, 0x73,
	0xad, 0x1a, 0x83, 0xaf, 0xa1, 0xab, 0xfd, 0x8c, 0x81, 0x6d, 0xd4, 0xea, 0x11, 0x85, 0xf7, 0x4f,
	0xdc, 0x18, 0xdf, 0x39, 0xf8, 0x0a, 0x3a, 0xca, 0x89, 0x78, 0xf9, 0x17, 0x9f, 0x87, 0xc1, 0x9f,
	0x17, 0xc7, 0x8f, 0xdf, 0x43, 0xff, 0xe0, 0x30, 0x7c, 0x70, 0xda, 0x77, 0x65, 0x90, 0xab, 0x36,
	0x53, 0x12, 0x47, 0x51, 0xe8, 0x09, 0xdb, 0x14, 0x75, 0x0f, 0xd9, 0x14, 0x96, 0x1d, 0x88, 0x83,
	0x9f, 0x61, 0x6c, 0xcf, 0x04, 0xad, 0x67, 0x70, 0x72, 0x90, 0x21, 0x69, 0x93, 0xd4, 0x19, 0x0f,
	0x73, 0xb1, 0x19, 0x1b, 0x43, 0xb5, 0x19, 0x9b, 0xa3, 0x24, 0xce, 0xd7, 0x9e, 0xfe, 0xdd, 0x3d,
	0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xd4, 0xa0, 0x23, 0x02, 0x05, 0x00, 0x00,
}
