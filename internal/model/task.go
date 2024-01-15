package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Task struct {
	Title    string `json:"title" gorm:"not null;size:128;comment:任务标题"`
	Desc     string `json:"desc" gorm:"not null;size:256;comment:任务公告"`
	IsUpdate *bool  `json:"is_update" gorm:"default false;comment:是否可更新"`
	Start    uint32 `json:"start" gorm:"not null;comment:开始时间"`
	End      uint32 `json:"end" gorm:"not null;comment:结束时间"`
	Config   string `json:"config" gorm:"not null;type:text;comment:任务配置"`
	types.BaseModel
}

// OneByID 通过id查询新闻信息
func (u *Task) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(u, id).Error
}

// Page 查询分页数据
func (u *Task) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Task, uint32, error) {
	var list []*Task
	var total int64

	db := ctx.DB().Model(u).
		Select("id", "title", "desc", "is_update", "start", "end", "created_at", "updated_at")

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
func (u *Task) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *Task) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除新闻信息
func (u *Task) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
