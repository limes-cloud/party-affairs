package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Resource struct {
	Title            string           `json:"title" gorm:"not null;size:128;comment:文档标题"`
	Desc             string           `json:"desc" gorm:"not null;size:256;comment:文档描述"`
	URL              string           `json:"url" gorm:"not null;size:256;comment:文档url"`
	DownloadCount    uint32           `json:"download_count" gorm:"default 0;comment:下载次数"`
	ClassifyID       uint32           `json:"classify_id" gorm:"not null;comment:新闻分类"`
	ResourceClassify ResourceClassify `json:"resource_classify" gorm:"foreignKey:classify_id"`
	types.BaseModel
}

// OneByID 通过id查询新闻信息
func (u *Resource) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Preload("ResourceClassify").First(u, id).Error
}

// Page 查询分页数据
func (u *Resource) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Resource, uint32, error) {
	var list []*Resource
	var total int64

	db := ctx.DB().Model(u).Preload("ResourceClassify")

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
func (u *Resource) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *Resource) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除新闻信息
func (u *Resource) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
