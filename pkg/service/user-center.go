package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	uc "github.com/limes-cloud/user-center/api/user/v1"
)

func NewUser(ctx context.Context, host string) (uc.ServiceClient, error) {
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(host))
	if err != nil {
		return nil, err
	}
	return uc.NewServiceClient(conn), nil
}
