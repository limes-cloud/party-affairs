package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	export "github.com/limes-cloud/resource/api/export/v1"
	file "github.com/limes-cloud/resource/api/file/v1"
	user "github.com/limes-cloud/user-center/api/user/v1"

	"party-affairs/api/errors"
)

const (
	Resource   = "Resource"
	UserCenter = "UserCenter"
)

func NewResourceExport(ctx context.Context) (export.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.Resource()
	}
	return export.NewServiceClient(conn), nil
}

func NewResourceFile(ctx context.Context) (file.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.Resource()
	}
	return file.NewServiceClient(conn), nil
}

func NewUser(ctx context.Context) (user.ServiceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(UserCenter)
	if err != nil {
		return nil, errors.Resource()
	}
	return user.NewServiceClient(conn), nil
}
