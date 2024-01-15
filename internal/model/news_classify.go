package model

import (
	"github.com/limes-cloud/kratosx"
)

type NewsClassify struct {
	Name   string `json:"name" gorm:"not null;size:128;comment:分类名称"`
	Weight uint32 `json:"weight" gorm:"comment:分类权重"`
	BaseModel
}

// All 查询全部数据
func (u *NewsClassify) All(ctx kratosx.Context) ([]*NewsClassify, error) {
	var list []*NewsClassify
	return list, ctx.DB().Model(u).Find(&list).Error
}

// Create 创建信息
func (u *NewsClassify) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *NewsClassify) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除信息
func (u *NewsClassify) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
