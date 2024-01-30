package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"

	v1 "party-affairs/api/v1"
)

func (s *Service) PageTask(ctx context.Context, in *v1.PageTaskRequest) (*v1.PageTaskReply, error) {
	return s.Task.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) GetTask(ctx context.Context, in *v1.GetTaskRequest) (*v1.Task, error) {
	return s.Task.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) AddTask(ctx context.Context, in *v1.AddTaskRequest) (*empty.Empty, error) {
	return s.Task.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateTask(ctx context.Context, in *v1.UpdateTaskRequest) (*empty.Empty, error) {
	return s.Task.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteTask(ctx context.Context, in *v1.DeleteTaskRequest) (*empty.Empty, error) {
	return s.Task.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) PageTaskValue(ctx context.Context, in *v1.PageTaskValueRequest) (*v1.PageTaskValueReply, error) {
	return s.TaskValue.Page(kratosx.MustContext(ctx), in)
}

func (s *Service) GetTaskValue(ctx context.Context, in *v1.GetTaskValueRequest) (*v1.TaskValue, error) {
	return s.TaskValue.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) AddTaskValue(ctx context.Context, in *v1.AddTaskValueRequest) (*empty.Empty, error) {
	return s.TaskValue.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateTaskValue(ctx context.Context, in *v1.UpdateTaskValueRequest) (*empty.Empty, error) {
	return s.TaskValue.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteTaskValue(ctx context.Context, in *v1.DeleteTaskValueRequest) (*empty.Empty, error) {
	return s.TaskValue.Delete(kratosx.MustContext(ctx), in)
}
