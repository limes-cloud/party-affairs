package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Auth struct {
	Platform string `json:"platform" gorm:"not null;size:32;comment:所属平台"`
	AuthID   string `json:"auth_id" gorm:"not null;size:128;comment:认证ID"`
	UnionID  string `json:"union_id" gorm:"default null;size:128;comment:联合ID"`
	UserID   uint32 `json:"user_id" gorm:"not null;comment:用户ID"`
	types.CreateModel
}

// OneByAuthId 通过uid查询用户信息
func (u *Auth) OneByAuthId(ctx kratosx.Context, authId string) error {
	return ctx.DB().First(u, "auth_id=?", authId).Error
}

// OneByUid 通过uid查询用户信息
func (u *Auth) OneByUid(ctx kratosx.Context, uid uint32) error {
	return ctx.DB().First(u, "user_id=?", uid).Error
}

// Create 创建用户信息
func (u *Auth) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *Auth) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除用户信息
func (u *Auth) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
