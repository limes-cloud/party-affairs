package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/service"
	"party-affairs/pkg/util"
)

type Banner struct {
	conf *config.Config
}

func NewBanner(conf *config.Config) *Banner {
	return &Banner{
		conf: conf,
	}
}

func (l *Banner) All(ctx kratosx.Context) (*v1.AllBannerReply, error) {
	banner := model.Banner{}
	list, err := banner.All(ctx)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	if len(list) == 0 {
		return &v1.AllBannerReply{List: make([]*v1.Banner, 0)}, nil
	}

	reply := v1.AllBannerReply{}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	resource, err := service.NewResource(ctx, l.conf.Service.Resource)
	if err != nil {
		return nil, v1.ResourceServiceError()
	}
	for ind, item := range reply.List {
		reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Src})
	}
	return &reply, nil
}

func (l *Banner) Add(ctx kratosx.Context, in *v1.AddBannerRequest) (*empty.Empty, error) {
	banner := model.Banner{}
	if err := util.Transform(in, &banner); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := banner.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Banner) Update(ctx kratosx.Context, in *v1.UpdateBannerRequest) (*empty.Empty, error) {
	banner := model.Banner{}
	if err := util.Transform(in, &banner); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := banner.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Banner) Delete(ctx kratosx.Context, in *v1.DeleteBannerRequest) (*empty.Empty, error) {
	banner := model.Banner{}
	if err := banner.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
