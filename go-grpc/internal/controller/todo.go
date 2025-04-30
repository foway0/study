package controller

import (
	"context"
	"github.com/foway0/study/go-grpc/api"
	"github.com/foway0/study/go-grpc/internal"
	"github.com/foway0/study/go-grpc/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type TodoServer struct {
	api.UnimplementedTodoServiceServer
}

func convert(todo repository.Todo) *api.Todo {
	return &api.Todo{
		Id:        todo.Id,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Body:      todo.Body,
		Status:    api.TodoStatus(api.TodoStatus_value[todo.Status]),
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *api.CreateTodoRequest) (*emptypb.Empty, error) {
	log.Println("Creating todo...")
	log.Println("Request:", req)

	_ctx := ctx.Value("_ctx").(*internal.ApplicationContext)
	now := time.Now().UnixMilli()
	repository.CreateTodo(_ctx.Mysql(), repository.Todo{
		UserId:    1,
		Title:     req.GetTitle(),
		Body:      req.GetBody(),
		CreatedAt: uint64(now),
		UpdatedAt: uint64(now),
	})

	return nil, nil
}

func (s *TodoServer) ListTodos(ctx context.Context, req *api.ListTodosRequest) (*api.ListTodosResponse, error) {
	log.Println("Listing todos...")
	log.Println("Request:", req)

	_ctx := ctx.Value("_ctx").(*internal.ApplicationContext)
	todos, err := repository.FindListTodo(_ctx.Mysql())
	if err != nil {
		log.Fatalf("failed to find todos: %v", err)
	}

	var todoList []*api.Todo
	for _, todo := range todos {
		log.Printf("Todo: %+v\n", todo)
		todoList = append(todoList, convert(todo))
	}

	return &api.ListTodosResponse{
		Todos: todoList,
	}, nil
}

func (s *TodoServer) GetTodoById(ctx context.Context, req *api.GetTodoByIdRequest) (*api.GetTodoByIdResponse, error) {
	log.Println("Getting todo by ID...")
	log.Println("Request:", req)
	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}

	_ctx := ctx.Value("_ctx").(*internal.ApplicationContext)
	todo, err := repository.FindOneTodo(_ctx.Mysql(), id)
	if err != nil {
		log.Fatalf("failed to find todos: %v", err)
	}

	if todo.Id == 0 {
		return nil, status.Errorf(codes.NotFound, "Todo with ID %d not found", id)
	}

	return &api.GetTodoByIdResponse{
		Todo: convert(todo),
	}, nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *api.UpdateTodoRequest) (*emptypb.Empty, error) {
	log.Println("Updating todo...")
	log.Println("Request:", req)
	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}

	_ctx := ctx.Value("_ctx").(*internal.ApplicationContext)
	now := time.Now().UnixMilli()
	repository.UpdateTodo(_ctx.Mysql(), req.GetId(), repository.Todo{
		Title:     req.GetTitle(),
		Body:      req.GetBody(),
		Status:    req.GetStatus().String(),
		UpdatedAt: uint64(now),
	})

	return nil, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *api.DeleteTodoRequest) (*emptypb.Empty, error) {
	log.Println("Deleting todo...")
	log.Println("Request:", req)
	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}

	_ctx := ctx.Value("_ctx").(*internal.ApplicationContext)
	repository.DeleteTodo(_ctx.Mysql(), id)

	return nil, nil
}
