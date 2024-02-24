package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"

	v1 "party-affairs/api/v1"
	"party-affairs/internal/biz"
	"party-affairs/internal/biz/types"
	"party-affairs/pkg/service"
)

func (s *Service) AllResourceClassify(ctx context.Context, _ *empty.Empty) (*v1.AllResourceClassifyReply, error) {
	list, err := s.resource.AllClassify(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	var reply v1.AllResourceClassifyReply
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddResourceClassify(ctx context.Context, in *v1.AddResourceClassifyRequest) (*empty.Empty, error) {
	var nc biz.ResourceClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.resource.AddClassify(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateResourceClassify(ctx context.Context, in *v1.UpdateResourceClassifyRequest) (*empty.Empty, error) {
	var nc biz.ResourceClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.resource.UpdateClassify(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteResourceClassify(ctx context.Context, in *v1.DeleteResourceClassifyRequest) (*empty.Empty, error) {
	return nil, s.resource.DeleteClassify(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) PageResourceContent(ctx context.Context, in *v1.PageResourceContentRequest) (*v1.PageResourceContentReply, error) {
	var req types.PageResourceContentRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, v1.TransformError()
	}

	list, total, err := s.resource.PageContent(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageResourceContentReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Url})
		}
	}

	return &reply, nil
}

func (s *Service) GetResourceContent(ctx context.Context, in *v1.GetResourceContentRequest) (*v1.ResourceContent, error) {
	nc, err := s.resource.GetContent(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.ResourceContent{}
	if err := copier.Copy(&reply, nc); err != nil {
		return nil, v1.TransformError()
	}

	resource, err := service.NewResource(ctx, s.conf.Service.Resource)
	if err == nil {
		reply.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.Url})
	}

	return &reply, nil
}

func (s *Service) AddResourceContent(ctx context.Context, in *v1.AddResourceContentRequest) (*empty.Empty, error) {
	var nc biz.ResourceContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.resource.AddContent(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateResourceContent(ctx context.Context, in *v1.UpdateResourceContentRequest) (*empty.Empty, error) {
	var nc biz.ResourceContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.resource.UpdateContent(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteResourceContent(ctx context.Context, in *v1.DeleteResourceContentRequest) (*empty.Empty, error) {
	return nil, s.resource.DeleteContent(kratosx.MustContext(ctx), in.Id)
}
