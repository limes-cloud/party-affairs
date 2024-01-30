package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
)

func (s *Service) AllNewsClassify(ctx context.Context, _ *empty.Empty) (*v1.AllNewsClassifyReply, error) {
	return s.NewsClassify.All(kratosx.MustContext(ctx))
}

func (s *Service) AddNewsClassify(ctx context.Context, in *v1.AddNewsClassifyRequest) (*empty.Empty, error) {
	return s.NewsClassify.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateNewsClassify(ctx context.Context, in *v1.UpdateNewsClassifyRequest) (*empty.Empty, error) {
	return s.NewsClassify.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteNewsClassify(ctx context.Context, in *v1.DeleteNewsClassifyRequest) (*empty.Empty, error) {
	return s.NewsClassify.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) PageNewsContent(ctx context.Context, in *v1.PageNewsRequest) (*v1.PageNewsReply, error) {
	return s.News.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) GetNewsContent(ctx context.Context, in *v1.GetNewsRequest) (*v1.News, error) {
	return s.News.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) AddNewsContent(ctx context.Context, in *v1.AddNewsRequest) (*empty.Empty, error) {
	return s.News.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateNewsContent(ctx context.Context, in *v1.UpdateNewsRequest) (*empty.Empty, error) {
	return s.News.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteNewsContent(ctx context.Context, in *v1.DeleteNewsRequest) (*empty.Empty, error) {
	return s.News.Delete(kratosx.MustContext(ctx), in)
}
