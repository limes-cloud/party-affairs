package service

import (
	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/biz"
	"party-affairs/internal/data"
)

// Service is a admin service.
type Service struct {
	v1.UnimplementedServiceServer
	conf     *config.Config
	notice   *biz.NoticeUseCase
	banner   *biz.BannerUseCase
	news     *biz.NewsUseCase
	resource *biz.ResourceUseCase
	task     *biz.TaskUseCase
	video    *biz.VideoUseCase
}

func New(conf *config.Config) *Service {
	return &Service{
		conf:     conf,
		notice:   biz.NewNoticeUseCase(conf, data.NewNoticeRepo()),
		banner:   biz.NewBannerUseCase(conf, data.NewBannerRepo()),
		news:     biz.NewNewsUseCase(conf, data.NewNewsRepo()),
		resource: biz.NewResourceUseCase(conf, data.NewResourceRepo()),
		task:     biz.NewTaskUseCase(conf, data.NewTaskRepo()),
		video:    biz.NewVideoUseCase(conf, data.NewVideoRepo()),
	}
}
