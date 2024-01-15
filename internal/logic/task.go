package logic

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"
	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/util"
)

type Task struct {
	conf *config.Config
}

func NewTask(conf *config.Config) *Task {
	return &Task{
		conf: conf,
	}
}

func (l *Task) Page(ctx kratosx.Context, in *v1.PageTaskRequest) (*v1.PageTaskReply, error) {
	task := model.Task{}
	list, total, err := task.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Title != nil {
				db = db.Where("title like ?", "%"+*in.Title+"%")
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageTaskReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (l *Task) Get(ctx kratosx.Context, in *v1.GetTaskRequest) (*v1.GetTaskReply, error) {
	task := model.Task{}
	if err := task.OneByID(ctx, in.Id); err != nil {
		return nil, v1.NotFoundError()
	}

	reply := v1.GetTaskReply{}
	if err := util.Transform(task, &reply.Task); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}
	return &reply, nil
}

func (l *Task) Add(ctx kratosx.Context, in *v1.AddTaskRequest) (*empty.Empty, error) {
	task := model.Task{}
	if err := util.Transform(in, &task); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := task.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Task) Update(ctx kratosx.Context, in *v1.UpdateTaskRequest) (*empty.Empty, error) {
	task := model.Task{}
	if err := util.Transform(in, &task); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := task.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *Task) Delete(ctx kratosx.Context, in *v1.DeleteTaskRequest) (*empty.Empty, error) {
	task := model.Task{}
	if err := task.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
