package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/util"
)

type ResourceClassify struct {
	conf *config.Config
}

func NewResourceClassify(conf *config.Config) *ResourceClassify {
	return &ResourceClassify{
		conf: conf,
	}
}

func (l *ResourceClassify) All(ctx kratosx.Context) (*v1.AllResourceClassifyReply, error) {
	nc := model.ResourceClassify{}
	list, err := nc.All(ctx)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.AllResourceClassifyReply{}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

func (l *ResourceClassify) Add(ctx kratosx.Context, in *v1.AddResourceClassifyRequest) (*empty.Empty, error) {
	user := model.ResourceClassify{}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *ResourceClassify) Update(ctx kratosx.Context, in *v1.UpdateResourceClassifyRequest) (*empty.Empty, error) {
	user := model.ResourceClassify{}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *ResourceClassify) Delete(ctx kratosx.Context, in *v1.DeleteResourceClassifyRequest) (*empty.Empty, error) {
	user := model.ResourceClassify{}
	if err := user.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
