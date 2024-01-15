package model

import (
	"github.com/limes-cloud/kratosx"
)

type News struct {
	Title        string       `json:"title" gorm:"not null;size:128;comment:新闻标题"`
	Unit         string       `json:"unit" gorm:"not null;size:128;comment:发布单位"`
	Cover        string       `json:"cover" gorm:"not null;size:128;comment:封面图片"`
	Desc         string       `json:"desc" gorm:"not null;size:128;comment:新闻描述"`
	Content      string       `json:"content" gorm:"not null;type:blob;comment:新闻内容"`
	Read         uint32       `json:"read" gorm:"default 0;comment:阅读人数"`
	ClassifyID   uint32       `json:"classify_id" gorm:"not null;comment:新闻分类"`
	NewsClassify NewsClassify `json:"news_classify" gorm:"foreignKey:classify_id"`
	BaseModel
}

// OneByID 通过id查询新闻信息
func (u *News) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Preload("NewsClassify").First(u, id).Error
}

// Page 查询分页数据
func (u *News) Page(ctx kratosx.Context, options *PageOptions) ([]*News, uint32, error) {
	var list []*News
	var total int64

	db := ctx.DB().Model(u).
		Select("id", "title", "unit", "cover", "desc", "read", "classify_id", "created_at").
		Preload("NewsClassify")

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
func (u *News) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *News) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除新闻信息
func (u *News) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
