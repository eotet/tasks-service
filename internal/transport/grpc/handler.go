package transport

import (
	"context"

	taskpb "github.com/eotet/project-protos/gen/go/task"
	userpb "github.com/eotet/project-protos/gen/go/user"
	"github.com/eotet/tasks-service/internal/task"
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
	panic("")
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	panic("")
}

func (h *Handler) ListTasks(ctx context.Context, req *emptypb.Empty) (*taskpb.ListTasksResponse, error) {
	panic("")
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	panic("")
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*emptypb.Empty, error) {
	panic("")
}
