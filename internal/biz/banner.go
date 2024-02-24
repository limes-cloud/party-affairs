package biz

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
)

type Banner struct {
	Title  string `json:"title" gorm:"not null;size:128;comment:轮播标题"`
	Src    string `json:"src" gorm:"not null;size:128;comment:轮播图"`
	Path   string `json:"path" gorm:"default null;size:256;comment:跳转链接"`
	Weight uint32 `json:"weight" gorm:"default 0;comment:轮播权重"`
	ktypes.BaseModel
}

type BannerRepo interface {
	All(ctx kratosx.Context) ([]*Banner, error)
	Create(ctx kratosx.Context, c *Banner) (uint32, error)
	Update(ctx kratosx.Context, c *Banner) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type BannerUseCase struct {
	config *config.Config
	repo   BannerRepo
}

func NewBannerUseCase(config *config.Config, repo BannerRepo) *BannerUseCase {
	return &BannerUseCase{config: config, repo: repo}
}

// All 获取全部轮播图
func (u *BannerUseCase) All(ctx kratosx.Context) ([]*Banner, error) {
	app, err := u.repo.All(ctx)
	if err != nil {
		return nil, v1.DatabaseError()
	}
	return app, nil
}

// Add 添加轮播图信息
func (u *BannerUseCase) Add(ctx kratosx.Context, app *Banner) (uint32, error) {
	id, err := u.repo.Create(ctx, app)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新轮播图信息
func (u *BannerUseCase) Update(ctx kratosx.Context, app *Banner) error {
	if err := u.repo.Update(ctx, app); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除轮播图信息
func (u *BannerUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
