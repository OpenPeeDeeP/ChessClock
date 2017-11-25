package main

import (
	"net"

	cc "github.com/OpenPeeDeeP/chessclock"
	"github.com/OpenPeeDeeP/chessclock/chessclock"
	"github.com/ianschenck/envflag"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var version string
var (
	daemonConString string
	maxFiles        int
)

func init() {
	envflag.StringVar(&daemonConString, "CCD_CONNECTION_STRING", "localhost:4242", "Connection string to the daemon")
	envflag.IntVar(&maxFiles, "CCD_MAX_FILES", 5, "Maximum number of log files")
}

func main() {
	envflag.Parse()
	lis, err := net.Listen("tcp", daemonConString)
	if err != nil {
		log.Error().Str("con", daemonConString).Msg("failed to listen")
	}
	grpcServer := grpc.NewServer()
	ccd := NewDaemon(cc.NewFileStore("OpenPeeDeeP", "ChessClock", maxFiles))
	chessclock.RegisterChessClockServer(grpcServer, ccd)
	grpcServer.Serve(lis)
}
