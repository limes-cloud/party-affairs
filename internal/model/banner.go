package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Banner struct {
	Title  string `json:"title" gorm:"not null;size:128;comment:轮播标题"`
	Src    string `json:"src" gorm:"not null;size:128;comment:轮播图"`
	App    string `json:"app" gorm:"default null;size:32;comment:所属应用"`
	Url    string `json:"url" gorm:"default null;size:128;comment:跳转链接"`
	Params string `json:"params" gorm:"default null;size:128;comment:跳转参数"`
	Weight uint32 `json:"weight" gorm:"default 0;comment:轮播权重"`
	types.BaseModel
}

// OneByID 通过id查询新闻信息
func (u *Banner) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(u, id).Error
}

// All 查询全部数据
func (u *Banner) All(ctx kratosx.Context) ([]*Banner, error) {
	var list []*Banner
	return list, ctx.DB().Model(u).Order("weight desc").Find(&list).Error
}

// Create 创建新闻信息
func (u *Banner) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *Banner) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除新闻信息
func (u *Banner) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
