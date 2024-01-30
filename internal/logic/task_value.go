package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/md"
	"party-affairs/pkg/util"
)

type TaskValue struct {
	conf *config.Config
}

func NewTaskValue(conf *config.Config) *TaskValue {
	return &TaskValue{
		conf: conf,
	}
}

func (l *TaskValue) Page(ctx kratosx.Context, in *v1.PageTaskValueRequest) (*v1.PageTaskValueReply, error) {
	task := model.TaskValue{}
	list, total, err := task.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.UserId != nil {
				db = db.Where("user_id = ?", in.UserId)
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageTaskValueReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (l *TaskValue) Get(ctx kratosx.Context, in *v1.GetTaskValueRequest) (*v1.TaskValue, error) {
	task := model.TaskValue{}
	if err := task.OneByID(ctx, in.TaskId, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	reply := v1.TaskValue{}
	if err := util.Transform(task, &reply.Value); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (l *TaskValue) Add(ctx kratosx.Context, in *v1.AddTaskValueRequest) (*empty.Empty, error) {
	task := model.TaskValue{
		TaskID: in.TaskId,
		UserID: md.UserId(ctx),
		Value:  in.Value,
	}

	if err := task.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *TaskValue) Update(ctx kratosx.Context, in *v1.UpdateTaskValueRequest) (*empty.Empty, error) {
	task := model.TaskValue{
		TaskID: in.TaskId,
		UserID: md.UserId(ctx),
		Value:  in.Value,
	}
	if err := task.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *TaskValue) Delete(ctx kratosx.Context, in *v1.DeleteTaskValueRequest) (*empty.Empty, error) {
	task := model.TaskValue{}
	if err := task.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
