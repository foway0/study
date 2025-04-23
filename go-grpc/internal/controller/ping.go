package controller

import (
	"context"
	"github.com/foway0/study/go-grpc/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingServer struct {
	api.UnimplementedPingServer
}

func (c *PingServer) Ping(_ context.Context, _ *emptypb.Empty) (*api.PingReply, error) {
	return &api.PingReply{
		Message: "Pong",
	}, nil
}
