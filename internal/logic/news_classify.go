package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/util"
)

type NewsClassify struct {
	conf *config.Config
}

func NewNewsClassify(conf *config.Config) *NewsClassify {
	return &NewsClassify{
		conf: conf,
	}
}

func (l *NewsClassify) All(ctx kratosx.Context) (*v1.AllNewsClassifyReply, error) {
	nc := model.NewsClassify{}
	list, err := nc.All(ctx)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.AllNewsClassifyReply{}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

func (l *NewsClassify) Add(ctx kratosx.Context, in *v1.AddNewsClassifyRequest) (*empty.Empty, error) {
	user := model.NewsClassify{}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *NewsClassify) Update(ctx kratosx.Context, in *v1.UpdateNewsClassifyRequest) (*empty.Empty, error) {
	user := model.NewsClassify{}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *NewsClassify) Delete(ctx kratosx.Context, in *v1.DeleteNewsClassifyRequest) (*empty.Empty, error) {
	user := model.NewsClassify{}
	if err := user.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
