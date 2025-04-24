# go-grpc

## Description
This is a simple gRPC server and client written in Go. It uses the `protoc` compiler to generate Go code from `.proto` files.
The server implements a simple calculator service that can perform addition and subtraction.

## Installation

```bash
# Install docker
brew install --cask docker
# Install Go
brew install go
# Install protoc (Protocol Buffers compiler)
brew install protobuf
# Install grpcurl (for testing)
brew install grpcurl
# Install protoc-gen-go (for generating Go code from .proto files)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# Install protoc-gen-go-grpc (for generating gRPC code from .proto files)
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## mod download

```bash
# Download the dependencies
go mod download
```

## Run the server

```bash
# Start the server
SERVICE_PORT=8080 go run cmd/go-grpc/main.go
```

## RUN by docker

```bash
make dressup
````

## Generate Go code from .proto files

protoc --go_out=./api --go_opt=paths=source_relative \
--go-grpc_out=./api --go-grpc_opt=paths=source_relative \
--proto_path=./api ./api/*.proto

## Test the server

```bash
# Test the server using grpcurl
grpcurl -plaintext localhost:8080 list
grpcurl -plaintext localhost:8080 api.Ping/Ping
```