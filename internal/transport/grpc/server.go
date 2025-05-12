package transport

import (
	"net"

	taskpb "github.com/eotet/project-protos/gen/go/task"
	userpb "github.com/eotet/project-protos/gen/go/user"
	"github.com/eotet/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, us userpb.UserServiceClient) error {
	listener, err := net.Listen("tcp", ":5502")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	handler := NewHandler(svc, us)

	taskpb.RegisterTaskServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
