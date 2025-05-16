dependency:
	@echo "Make sure you have the following dependencies installed:"
	@command -v go
	@command -v protoc
	@command -v docker
	@command -v docker-compose
	@command -v pre-commit
	@echo "Install the following dependencies:"
	@echo "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
	@echo "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
	@echo "docker network create my_network"

dressup:
	docker-compose -f ./build/docker-compose.yaml up --build
