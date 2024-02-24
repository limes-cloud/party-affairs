package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/user-center/api/auth"
	userV1 "github.com/limes-cloud/user-center/api/user/v1"

	v1 "party-affairs/api/v1"
	"party-affairs/internal/biz"
	"party-affairs/internal/biz/types"
	"party-affairs/pkg/service"
)

func (s *Service) PageTask(ctx context.Context, in *v1.PageTaskRequest) (*v1.PageTaskReply, error) {
	list, total, err := s.task.Page(kratosx.MustContext(ctx), &types.PageTaskRequest{
		Page:      in.Page,
		PageSize:  in.PageSize,
		Title:     in.Title,
		NotFinish: in.NotFinish,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.PageTaskReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) GetTask(ctx context.Context, in *v1.GetTaskRequest) (*v1.Task, error) {
	task, err := s.task.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	reply := v1.Task{}
	if err := copier.Copy(&reply, task); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddTask(ctx context.Context, in *v1.AddTaskRequest) (*empty.Empty, error) {
	var task biz.Task
	if err := copier.Copy(&task, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.task.Add(kratosx.MustContext(ctx), &task)
	return nil, err
}

func (s *Service) UpdateTask(ctx context.Context, in *v1.UpdateTaskRequest) (*empty.Empty, error) {
	var task biz.Task
	if err := copier.Copy(&task, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.task.Update(kratosx.MustContext(ctx), &task)
}

func (s *Service) DeleteTask(ctx context.Context, in *v1.DeleteTaskRequest) (*empty.Empty, error) {
	return nil, s.task.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) PageTaskValue(ctx context.Context, in *v1.PageTaskValueRequest) (*v1.PageTaskValueReply, error) {
	list, total, err := s.task.PageValue(kratosx.MustContext(ctx), &types.PageTaskValueRequest{
		Page:     in.Page,
		PageSize: in.PageSize,
		TaskID:   in.TaskId,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.PageTaskValueReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}

	user, err := service.NewUser(ctx, s.conf.Service.UserCenter)
	if err == nil {
		for ind, item := range reply.List {
			reply.List[ind].User, _ = user.GetBaseUser(ctx, &userV1.GetBaseUserRequest{Id: item.UserId})
		}
	}
	return &reply, nil
}

func (s *Service) GetTaskValue(ctx context.Context, in *v1.GetTaskValueRequest) (*v1.TaskValue, error) {
	task, err := s.task.GetValue(kratosx.MustContext(ctx), in.TaskId, in.UserId)
	if err != nil {
		return nil, err
	}
	reply := v1.TaskValue{}
	if err := copier.Copy(&reply, task); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) GetCurTaskValue(ctx context.Context, in *v1.GetCurTaskValueRequest) (*v1.TaskValue, error) {
	kCtx := kratosx.MustContext(ctx)

	md, err := auth.Get(kCtx)
	if err != nil {
		return nil, v1.AuthInfoError()
	}
	task, err := s.task.GetValue(kratosx.MustContext(ctx), in.TaskId, md.UserID)
	if err != nil {
		return nil, err
	}
	reply := v1.TaskValue{}
	if err := copier.Copy(&reply, task); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

func (s *Service) AddTaskValue(ctx context.Context, in *v1.AddTaskValueRequest) (*empty.Empty, error) {
	var task biz.TaskValue
	if err := copier.Copy(&task, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.task.AddValue(kratosx.MustContext(ctx), &task)
	return nil, err
}

func (s *Service) UpdateTaskValue(ctx context.Context, in *v1.UpdateTaskValueRequest) (*empty.Empty, error) {
	var task biz.TaskValue
	if err := copier.Copy(&task, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.task.UpdateValue(kratosx.MustContext(ctx), &task)
}

func (s *Service) DeleteTaskValue(ctx context.Context, in *v1.DeleteTaskValueRequest) (*empty.Empty, error) {
	return nil, s.task.DeleteValue(kratosx.MustContext(ctx), in.Id)
}
