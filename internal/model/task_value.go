package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type TaskValue struct {
	TaskID uint32 `json:"task_id"  gorm:"not null;comment:任务id"`
	UserID uint32 `json:"user_id" gorm:"not null;comment:用户id"`
	Value  string `json:"value"  gorm:"not null;type:text;comment:数据值"`
	types.BaseModel
}

// OneByID 通过id查询新闻信息
func (u *TaskValue) OneByID(ctx kratosx.Context, taskId uint32, userId uint32) error {
	return ctx.DB().Where("task_id=? and user_id=?", taskId, userId).First(u).Error
}

// Page 查询分页数据
func (u *TaskValue) Page(ctx kratosx.Context, options *types.PageOptions) ([]*TaskValue, uint32, error) {
	var list []*TaskValue
	var total int64

	db := ctx.DB().Model(u)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Create 创建新闻信息
func (u *TaskValue) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *TaskValue) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Where("task_id=? and user_id=?", u.TaskID, u.UserID).Updates(&u).Error
}

// DeleteByID 通过id删除新闻信息
func (u *TaskValue) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
