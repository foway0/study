package internal

import (
	"fmt"
	"github.com/foway0/study/go-grpc/api"
	"github.com/foway0/study/go-grpc/internal/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Boot initializes the gRPC server and starts listening for incoming connections.
func Boot() {
	log.Println("Starting gRPC server...")
	_config := GetConfig()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", _config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	api.RegisterPingServer(server, &controller.PingServer{})

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
