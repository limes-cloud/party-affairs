package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
)

func (s *Service) AllResourceClassify(ctx context.Context, _ *empty.Empty) (*v1.AllResourceClassifyReply, error) {
	return s.ResourceClassify.All(kratosx.MustContext(ctx))
}

func (s *Service) AddResourceClassify(ctx context.Context, in *v1.AddResourceClassifyRequest) (*empty.Empty, error) {
	return s.ResourceClassify.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateResourceClassify(ctx context.Context, in *v1.UpdateResourceClassifyRequest) (*empty.Empty, error) {
	return s.ResourceClassify.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteResourceClassify(ctx context.Context, in *v1.DeleteResourceClassifyRequest) (*empty.Empty, error) {
	return s.ResourceClassify.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) PageResource(ctx context.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	return s.Resource.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) GetResource(ctx context.Context, in *v1.GetResourceRequest) (*v1.GetResourceReply, error) {
	return s.Resource.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) AddResource(ctx context.Context, in *v1.AddResourceRequest) (*empty.Empty, error) {
	return s.Resource.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateResource(ctx context.Context, in *v1.UpdateResourceRequest) (*empty.Empty, error) {
	return s.Resource.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteResource(ctx context.Context, in *v1.DeleteResourceRequest) (*empty.Empty, error) {
	return s.Resource.Delete(kratosx.MustContext(ctx), in)
}
