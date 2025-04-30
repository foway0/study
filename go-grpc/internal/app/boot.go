package app

import (
	"context"
	"fmt"
	"github.com/foway0/study/go-grpc/api"
	"github.com/foway0/study/go-grpc/internal"
	"github.com/foway0/study/go-grpc/internal/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Boot initializes the gRPC server and starts listening for incoming connections.
func Boot() {
	log.Println("Initializing Context...")
	appCtx := internal.NewApplicationContext()
	log.Println("Context initialized.")

	log.Println("Starting gRPC server...")
	log.Println("Service Port: ", appCtx.Config().Port)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", appCtx.Config().Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("initializing gRPC middleware...")
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				ctx = context.WithValue(ctx, "_ctx", appCtx)
				return handler(ctx, req)
			},
		),
	)

	api.RegisterPingServer(server, &controller.PingServer{})
	api.RegisterTodoServiceServer(server, &controller.TodoServer{})

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
