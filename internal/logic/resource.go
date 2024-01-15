package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	resourceV1 "github.com/limes-cloud/resource/api/v1"
	"gorm.io/gorm"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/service"
	"party-affairs/pkg/util"
)

type Resource struct {
	conf *config.Config
}

func NewResource(conf *config.Config) *Resource {
	return &Resource{
		conf: conf,
	}
}

func (l *Resource) Page(ctx kratosx.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	rs := model.Resource{}
	list, total, err := rs.Page(ctx, &model.PageOptions{
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

	reply := v1.PageResourceReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	resource, err := service.NewResource(ctx, l.conf.Service.Resource)
	if err != nil {
		return nil, v1.ResourceServiceError()
	}
	for ind, item := range reply.List {
		reply.List[ind].Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: item.Url})
	}
	return &reply, nil
}

func (l *Resource) Get(ctx kratosx.Context, in *v1.GetResourceRequest) (*v1.GetResourceReply, error) {
	rs := model.Resource{}
	if err := rs.OneByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	reply := v1.GetResourceReply{}
	if err := util.Transform(rs, &reply.Resource); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	resource, err := service.NewResource(ctx, l.conf.Service.Resource)
	if err != nil {
		return nil, v1.ResourceServiceError()
	}
	reply.Resource.Resource, _ = resource.GetFileBySha(ctx, &resourceV1.GetFileByShaRequest{Sha: reply.Resource.Url})
	return &reply, nil
}

func (l *Resource) Add(ctx kratosx.Context, in *v1.AddResourceRequest) (*empty.Empty, error) {
	rs := model.Resource{}
	if err := util.Transform(in, &rs); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := rs.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Resource) Update(ctx kratosx.Context, in *v1.UpdateResourceRequest) (*empty.Empty, error) {
	rs := model.Resource{}
	if err := util.Transform(in, &rs); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := rs.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Resource) Delete(ctx kratosx.Context, in *v1.DeleteResourceRequest) (*empty.Empty, error) {
	rs := model.Resource{}
	if err := rs.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
