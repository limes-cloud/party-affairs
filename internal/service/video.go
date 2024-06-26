package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/file/v1"

	"party-affairs/api/errors"
	v1 "party-affairs/api/v1"
	"party-affairs/internal/biz"
	"party-affairs/internal/biz/types"
	"party-affairs/internal/pkg/service"
)

func (s *Service) PageVideoClassify(ctx context.Context, in *v1.PageVideoClassifyRequest) (*v1.PageVideoClassifyReply, error) {
	list, total, err := s.video.PageClassify(kratosx.MustContext(ctx), &types.PageVideoClassifyRequest{
		Page:     in.Page,
		PageSize: in.PageSize,
		Name:     in.Name,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.PageVideoClassifyReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}
	resource, err := service.NewResourceFile(ctx)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Cover})
		}
	}
	return &reply, nil
}

func (s *Service) AddVideoClassify(ctx context.Context, in *v1.AddVideoClassifyRequest) (*empty.Empty, error) {
	var nc biz.VideoClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.Transform()
	}
	_, err := s.video.AddClassify(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateVideoClassify(ctx context.Context, in *v1.UpdateVideoClassifyRequest) (*empty.Empty, error) {
	var nc biz.VideoClassify
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.Transform()
	}
	return nil, s.video.UpdateClassify(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteVideoClassify(ctx context.Context, in *v1.DeleteVideoClassifyRequest) (*empty.Empty, error) {
	return nil, s.video.DeleteClassify(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) PageVideoContent(ctx context.Context, in *v1.PageVideoContentRequest) (*v1.PageVideoContentReply, error) {
	var req types.PageVideoContentRequest
	if err := copier.Copy(&req, in); err != nil {
		return nil, errors.Transform()
	}

	list, total, err := s.video.PageContent(kratosx.MustContext(ctx), &req)
	if err != nil {
		return nil, err
	}

	reply := v1.PageVideoContentReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.Transform()
	}

	video, err := service.NewResourceFile(ctx)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].Resource, _ = video.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Url})
		}
	}

	return &reply, nil
}

func (s *Service) GetVideoContent(ctx context.Context, in *v1.GetVideoContentRequest) (*v1.VideoContent, error) {
	nc, err := s.video.GetContent(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.VideoContent{}
	if err := copier.Copy(&reply, nc); err != nil {
		return nil, errors.Transform()
	}

	video, err := service.NewResourceFile(ctx)
	if err == nil {
		reply.Resource, _ = video.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.Url})
	}

	return &reply, nil
}

func (s *Service) AddVideoContent(ctx context.Context, in *v1.AddVideoContentRequest) (*empty.Empty, error) {
	var nc biz.VideoContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.Transform()
	}
	_, err := s.video.AddContent(kratosx.MustContext(ctx), &nc)
	return nil, err
}

func (s *Service) UpdateVideoContent(ctx context.Context, in *v1.UpdateVideoContentRequest) (*empty.Empty, error) {
	var nc biz.VideoContent
	if err := copier.Copy(&nc, in); err != nil {
		return nil, errors.Transform()
	}
	return nil, s.video.UpdateContent(kratosx.MustContext(ctx), &nc)
}

func (s *Service) DeleteVideoContent(ctx context.Context, in *v1.DeleteVideoContentRequest) (*empty.Empty, error) {
	return nil, s.video.DeleteContent(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) UpdateUserVideoProcess(ctx context.Context, in *v1.UpdateUserVideoProcessRequest) (*empty.Empty, error) {
	return nil, s.video.UpdateUserVideoProcess(kratosx.MustContext(ctx), in.VideoId, in.Time)
}
