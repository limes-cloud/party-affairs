package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	resource "github.com/limes-cloud/resource/api/v1"
	user "github.com/limes-cloud/user-center/api/user/v1"

	v1 "party-affairs/api/v1"
)

const (
	Resource   = "Resource"
	UserCenter = "UserCenter"
)

func NewResource(ctx context.Context) (resource.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, v1.ResourceError()
	}
	return resource.NewServiceClient(conn), nil
}

func NewUser(ctx context.Context) (user.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(UserCenter)
	if err != nil {
		return nil, v1.ResourceError()
	}
	return user.NewServiceClient(conn), nil
}
