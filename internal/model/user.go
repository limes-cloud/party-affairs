package model

import (
	"github.com/limes-cloud/kratosx"
)

type User struct {
	Avatar    string `json:"avatar" gorm:"not null;size:128;comment:用户头像"`
	Name      string `json:"name" gorm:"not null;size:32;comment:用户姓名"`
	Nickname  string `json:"nickname" gorm:"not null;size:64;comment:用户昵称"`
	Gender    string `json:"gender" gorm:"not null;comment:用户性别"`
	Phone     string `json:"phone" gorm:"not null;size:11;comment:用户电话"`
	Email     string `json:"email" gorm:"not null;type:varbinary(128);comment:用户邮箱"`
	Status    *bool  `json:"status,omitempty" gorm:"not null;comment:用户状态"`
	Token     string `json:"-" gorm:"default null;type:text;comment:登录token"`
	LastLogin uint32 `json:"last_login" gorm:"comment:最后登陆时间"`
	BaseModel
}

// OneByID 通过id查询用户信息
func (u *User) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(u, id).Error
}

// OneByPhone 通过phone查询用户信息
func (u *User) OneByPhone(ctx kratosx.Context, phone string) error {
	return ctx.DB().First(u, "phone=?", phone).Error
}

// OneByEmail 通过email查询用户信息
func (u *User) OneByEmail(ctx kratosx.Context, email string) error {
	return ctx.DB().First(u, "email=?", email).Error
}

// Page 查询分页数据
func (u *User) Page(ctx kratosx.Context, options *PageOptions) ([]*User, uint32, error) {
	var list []*User
	var total int64

	db := ctx.DB().Model(u) // .Preload("Role").Preload("Department")
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// Create 创建用户信息
func (u *User) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(u).Error
}

// Update 保存当前信息
func (u *User) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(u).Updates(&u).Error
}

// DeleteByID 通过id删除用户信息
func (u *User) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(u, id).Error
}
