package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
)

func (s *Service) AllBanner(ctx context.Context, _ *empty.Empty) (*v1.AllBannerReply, error) {
	return s.Banner.All(kratosx.MustContext(ctx))
}

func (s *Service) AddBanner(ctx context.Context, in *v1.AddBannerRequest) (*empty.Empty, error) {
	return s.Banner.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateBanner(ctx context.Context, in *v1.UpdateBannerRequest) (*empty.Empty, error) {
	return s.Banner.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteBanner(ctx context.Context, in *v1.DeleteBannerRequest) (*empty.Empty, error) {
	return s.Banner.Delete(kratosx.MustContext(ctx), in)
}
