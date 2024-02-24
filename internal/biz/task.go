package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/user-center/api/auth"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/biz/types"
)

type Task struct {
	Title     string     `json:"title" gorm:"not null;size:128;comment:任务标题"`
	Desc      string     `json:"desc" gorm:"not null;size:256;comment:任务公告"`
	IsUpdate  *bool      `json:"is_update" gorm:"default false;comment:是否可更新"`
	Start     uint32     `json:"start" gorm:"not null;comment:开始时间"`
	End       uint32     `json:"end" gorm:"not null;comment:结束时间"`
	Config    string     `json:"config" gorm:"not null;type:text;comment:任务配置"`
	TaskValue *TaskValue `json:"task_value"  gorm:"constraint:onDelete:cascade"`
	ktypes.BaseModel
}

type TaskValue struct {
	TaskID uint32 `json:"task_id"  gorm:"uniqueIndex:tu;not null;comment:任务id"`
	UserID uint32 `json:"user_id" gorm:"uniqueIndex:tu;not null;comment:用户id"`
	Value  string `json:"value"  gorm:"not null;type:text;comment:数据值"`
	ktypes.BaseModel
}

type TaskRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Task, error)
	Page(ctx kratosx.Context, in *types.PageTaskRequest) ([]*Task, uint32, error)
	Create(ctx kratosx.Context, c *Task) (uint32, error)
	Update(ctx kratosx.Context, c *Task) error
	Delete(ctx kratosx.Context, uint322 uint32) error

	GetValue(ctx kratosx.Context, taskId, userId uint32) (*TaskValue, error)
	PageValue(ctx kratosx.Context, in *types.PageTaskValueRequest) ([]*TaskValue, uint32, error)
	CreateValue(ctx kratosx.Context, c *TaskValue) (uint32, error)
	UpdateValue(ctx kratosx.Context, c *TaskValue) error
	DeleteValue(ctx kratosx.Context, uint322 uint32) error
}

type TaskUseCase struct {
	config *config.Config
	repo   TaskRepo
}

func NewTaskUseCase(config *config.Config, repo TaskRepo) *TaskUseCase {
	return &TaskUseCase{config: config, repo: repo}
}

// Get 获取指定的公告
func (u *TaskUseCase) Get(ctx kratosx.Context, id uint32) (*Task, error) {
	task, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.DatabaseError()
	}
	return task, nil
}

// Page 获取分页任务
func (u *TaskUseCase) Page(ctx kratosx.Context, req *types.PageTaskRequest) ([]*Task, uint32, error) {
	md, err := auth.Get(ctx)
	if err == nil && md.UserID != 0 {
		req.UserID = &md.UserID
	}
	task, total, err := u.repo.Page(ctx, req)
	if err != nil {
		return nil, total, v1.DatabaseError()
	}
	return task, total, nil
}

// Add 添加任务信息
func (u *TaskUseCase) Add(ctx kratosx.Context, task *Task) (uint32, error) {
	id, err := u.repo.Create(ctx, task)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新任务信息
func (u *TaskUseCase) Update(ctx kratosx.Context, task *Task) error {
	if err := u.repo.Update(ctx, task); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除任务信息
func (u *TaskUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// GetValue 获取指定的值
func (u *TaskUseCase) GetValue(ctx kratosx.Context, taskId, userId uint32) (*TaskValue, error) {
	task, _ := u.repo.GetValue(ctx, taskId, userId)
	return task, nil
}

// PageValue 获取分页任务
func (u *TaskUseCase) PageValue(ctx kratosx.Context, in *types.PageTaskValueRequest) ([]*TaskValue, uint32, error) {
	task, total, err := u.repo.PageValue(ctx, in)
	if err != nil {
		return nil, total, v1.DatabaseError()
	}
	return task, total, nil
}

// AddValue 添加任务信息
func (u *TaskUseCase) AddValue(ctx kratosx.Context, task *TaskValue) (uint32, error) {
	md, err := auth.Get(ctx)
	if err != nil {
		return 0, v1.AuthInfoError()
	}
	task.UserID = md.UserID

	id, err := u.repo.CreateValue(ctx, task)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// UpdateValue 更新任务信息
func (u *TaskUseCase) UpdateValue(ctx kratosx.Context, task *TaskValue) error {
	md, err := auth.Get(ctx)
	if err != nil {
		return v1.AuthInfoError()
	}
	task.UserID = md.UserID

	if err := u.repo.UpdateValue(ctx, task); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// DeleteValue 删除任务信息
func (u *TaskUseCase) DeleteValue(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteValue(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
