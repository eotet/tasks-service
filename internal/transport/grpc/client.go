package transport

import (
	userpb "github.com/eotet/project-protos/gen/go/user"
	"google.golang.org/grpc"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := userpb.NewUserServiceClient(conn)
	return client, conn, nil
}
