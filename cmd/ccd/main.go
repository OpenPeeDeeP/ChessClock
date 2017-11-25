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
)

func init() {
	envflag.StringVar(&daemonConString, "CCD_CONNECTION_STRING", "localhost:4242", "Connection string to the daemon")
}

func main() {
	lis, err := net.Listen("tcp", daemonConString)
	if err != nil {
		log.Error().Str("con", daemonConString).Msg("failed to listen")
	}
	grpcServer := grpc.NewServer()
	ccd := NewDaemon(&cc.NilStore{})
	chessclock.RegisterChessClockServer(grpcServer, ccd)
	grpcServer.Serve(lis)
}
