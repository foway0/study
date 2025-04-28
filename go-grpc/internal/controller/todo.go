package controller

import (
	"context"
	"github.com/foway0/study/go-grpc/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type TodoServer struct {
	api.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(_ context.Context, req *api.CreateTodoRequest) (*api.CreateTodoResponse, error) {
	log.Println("Creating todo...")

	todo := &api.CreateTodoResponse{
		Todo: &api.Todo{
			Id:        1,
			Title:     req.GetTitle(),
			Body:      req.GetBody(),
			Status:    api.TodoStatus_in_progress,
			CreatedAt: uint64(time.Now().UnixMilli()),
			UpdatedAt: uint64(time.Now().UnixMilli()),
		},
	}

	return todo, nil
}

func (s *TodoServer) ListTodos(_ context.Context, req *api.ListTodosRequest) (*api.ListTodosResponse, error) {
	log.Println("Listing todos...")
	log.Println("Request:", req)

	todos := &api.ListTodosResponse{
		Todos: []*api.Todo{
			{
				Id:        1,
				Title:     "Todo 1",
				Body:      "Todo 1 body",
				Status:    api.TodoStatus_todo,
				CreatedAt: uint64(time.Now().UnixMilli()),
				UpdatedAt: uint64(time.Now().UnixMilli()),
			},
			{
				Id:        2,
				Title:     "Todo 2",
				Body:      "Todo 2 body",
				Status:    api.TodoStatus_in_progress,
				CreatedAt: uint64(time.Now().UnixMilli()),
				UpdatedAt: uint64(time.Now().UnixMilli()),
			},
			{
				Id:        3,
				Title:     "Todo 3",
				Body:      "Todo 3 body",
				Status:    api.TodoStatus_done,
				CreatedAt: uint64(time.Now().UnixMilli()),
				UpdatedAt: uint64(time.Now().UnixMilli()),
			},
		},
	}

	return todos, nil
}

func (s *TodoServer) GetTodoById(_ context.Context, req *api.GetTodoByIdRequest) (*api.GetTodoByIdResponse, error) {
	log.Println("Getting todo by ID...")
	log.Println("Request:", req)

	todo := &api.GetTodoByIdResponse{
		Todo: &api.Todo{
			Id:        req.GetId(),
			Title:     "Todo 1",
			Body:      "Todo 1 body",
			Status:    api.TodoStatus_in_progress,
			CreatedAt: uint64(time.Now().UnixMilli()),
			UpdatedAt: uint64(time.Now().UnixMilli()),
		},
	}

	return todo, nil
}

func (s *TodoServer) UpdateTodo(_ context.Context, req *api.UpdateTodoRequest) (*api.UpdateTodoResponse, error) {
	log.Println("Updating todo...")
	log.Println("Request:", req)

	todo := &api.UpdateTodoResponse{
		Todo: &api.Todo{
			Id:        req.GetId(),
			Title:     req.GetTitle(),
			Body:      req.GetBody(),
			Status:    req.GetStatus(),
			CreatedAt: uint64(time.Now().UnixMilli()),
			UpdatedAt: uint64(time.Now().UnixMilli()),
		},
	}

	return todo, nil
}

func (s *TodoServer) DeleteTodo(_ context.Context, req *api.DeleteTodoRequest) (*emptypb.Empty, error) {
	log.Println("Deleting todo...")
	log.Println("Request:", req)

	return nil, nil
}
