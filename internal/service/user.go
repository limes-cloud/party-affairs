package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "party-affairs/api/v1"
)

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.Auth.Login(kratosx.MustContext(ctx), in)
}

func (s *Service) SendBindEmail(ctx context.Context, in *v1.SendBindEmailRequest) (*v1.SendBindEmailReply, error) {
	return s.Auth.SendBindEmail(kratosx.MustContext(ctx), in)
}

func (s *Service) Bind(ctx context.Context, in *v1.BindRequest) (*v1.BindReply, error) {
	return s.Auth.Bind(kratosx.MustContext(ctx), in)
}

func (s *Service) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	return s.Auth.RefreshToken(kratosx.MustContext(ctx))
}

func (s *Service) PageUser(ctx context.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	return s.User.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) ChangeUserStatus(ctx context.Context, in *v1.ChangeUserStatusRequest) (*empty.Empty, error) {
	return s.User.ChangeStatus(kratosx.MustContext(ctx), in)
}

func (s *Service) CurrentUser(ctx context.Context, _ *empty.Empty) (*v1.User, error) {
	return s.User.Current(kratosx.MustContext(ctx))
}

func (s *Service) AddUser(ctx context.Context, in *v1.AddUserRequest) (*empty.Empty, error) {
	return s.User.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*empty.Empty, error) {
	return s.User.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*empty.Empty, error) {
	return s.User.Delete(kratosx.MustContext(ctx), in)
}
