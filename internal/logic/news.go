package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	resourceV1 "github.com/limes-cloud/resource/api/v1"
	"gorm.io/gorm"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/service"
	"party-affairs/pkg/util"
)

type News struct {
	conf *config.Config
}

func NewNews(conf *config.Config) *News {
	return &News{
		conf: conf,
	}
}

func (l *News) Page(ctx kratosx.Context, in *v1.PageNewsRequest) (*v1.PageNewsReply, error) {
	news := model.News{}
	list, total, err := news.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.ClassifyId != nil && *in.ClassifyId != 0 {
				db = db.Where("classify_id=?", *in.ClassifyId)
			}

			if in.Title != nil {
				db = db.Where("title like ?", "%"+*in.Title+"%")
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageNewsReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	resource, err := service.NewResource(ctx, l.conf.Service.Resource)
	if err != nil {
		return nil, v1.ResourceServiceError()
	}
	for ind, item := range reply.List {
		reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Cover})
	}
	return &reply, nil
}

func (l *News) Get(ctx kratosx.Context, in *v1.GetNewsRequest) (*v1.GetNewsReply, error) {
	news := model.News{}
	if err := news.OneByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	reply := v1.GetNewsReply{}
	if err := util.Transform(news, &reply.News); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	resource, err := service.NewResource(ctx, l.conf.Service.Resource)
	if err != nil {
		return nil, v1.ResourceServiceError()
	}
	reply.News.Resource, err = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.News.Cover})
	return &reply, nil
}

func (l *News) Add(ctx kratosx.Context, in *v1.AddNewsRequest) (*empty.Empty, error) {
	news := model.News{}
	if err := util.Transform(in, &news); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := news.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *News) Update(ctx kratosx.Context, in *v1.UpdateNewsRequest) (*empty.Empty, error) {
	news := model.News{}
	if err := util.Transform(in, &news); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := news.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *News) Delete(ctx kratosx.Context, in *v1.DeleteNewsRequest) (*empty.Empty, error) {
	news := model.News{}
	if err := news.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
