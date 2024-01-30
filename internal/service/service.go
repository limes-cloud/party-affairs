package service

import (
	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/logic"
)

// Service is a admin service.
type Service struct {
	v1.UnimplementedServiceServer
	News             *logic.News
	NewsClassify     *logic.NewsClassify
	Resource         *logic.Resource
	ResourceClassify *logic.ResourceClassify
	Banner           *logic.Banner
	Task             *logic.Task
	TaskValue        *logic.TaskValue
}

func New(conf *config.Config) *Service {
	return &Service{
		News:             logic.NewNews(conf),
		NewsClassify:     logic.NewNewsClassify(conf),
		Resource:         logic.NewResource(conf),
		ResourceClassify: logic.NewResourceClassify(conf),
		Banner:           logic.NewBanner(conf),
		Task:             logic.NewTask(conf),
		TaskValue:        logic.NewTaskValue(conf),
	}
}
