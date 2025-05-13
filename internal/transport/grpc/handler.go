package transport

import (
	"context"
	"errors"
	"fmt"

	taskpb "github.com/eotet/project-protos/gen/go/task"
	userpb "github.com/eotet/project-protos/gen/go/user"
	errs "github.com/eotet/tasks-service/internal/errors"
	"github.com/eotet/tasks-service/internal/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, us userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: us}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	task, err := h.svc.CreateTask(task.Task{UserId: uint(req.UserId), Text: req.Title})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(task.ID),
			UserId: uint32(task.UserId),
			Title:  task.Text,
			IsDone: task.IsDone,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	task, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		if errors.Is(err, errs.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get task: %v", err)
	}

	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(task.ID),
			UserId: uint32(task.UserId),
			Title:  task.Text,
			IsDone: task.IsDone,
		},
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *emptypb.Empty) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get tasks: %v", err)
	}

	var response taskpb.ListTasksResponse
	for i := range tasks {
		response.Tasks = append(response.Tasks, &taskpb.Task{
			Id:     uint32(tasks[i].ID),
			UserId: uint32(tasks[i].UserId),
			Title:  tasks[i].Text,
			IsDone: tasks[i].IsDone,
		})
	}

	return &response, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	updatedTask, err := h.svc.UpdateTaskByID(req.Id, task.UpdateTaskRequest{
		Text:   &req.Title,
		IsDone: &req.IsDone,
	})
	if err != nil {
		if errors.Is(err, errs.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to updated task: %v", err)
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     uint32(updatedTask.ID),
			UserId: uint32(updatedTask.UserId),
			Title:  updatedTask.Text,
			IsDone: updatedTask.IsDone,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*emptypb.Empty, error) {
	if err := h.svc.DeleteTaskByID(req.Id); err != nil {
		if errors.Is(err, errs.ErrTaskNotFound) {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete task: %v", err)
	}

	return &emptypb.Empty{}, nil
}
