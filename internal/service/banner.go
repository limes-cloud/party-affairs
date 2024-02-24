package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"

	v1 "party-affairs/api/v1"
	"party-affairs/internal/biz"
	"party-affairs/pkg/service"
)

func (s *Service) AllBanner(ctx context.Context, _ *empty.Empty) (*v1.AllBannerReply, error) {
	list, err := s.banner.All(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}

	var reply v1.AllBannerReply
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Src})
		}
	}
	return &reply, nil
}

func (s *Service) AddBanner(ctx context.Context, in *v1.AddBannerRequest) (*empty.Empty, error) {
	var banner biz.Banner
	if err := copier.Copy(&banner, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.banner.Add(kratosx.MustContext(ctx), &banner)
	return nil, err
}

func (s *Service) UpdateBanner(ctx context.Context, in *v1.UpdateBannerRequest) (*empty.Empty, error) {
	var banner biz.Banner
	if err := copier.Copy(&banner, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.banner.Update(kratosx.MustContext(ctx), &banner)
}

func (s *Service) DeleteBanner(ctx context.Context, in *v1.DeleteBannerRequest) (*empty.Empty, error) {
	return nil, s.banner.Delete(kratosx.MustContext(ctx), in.Id)
}
