package logic

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "party-affairs/api/v1"
	"party-affairs/config"
	"party-affairs/internal/model"
	"party-affairs/pkg/md"
	"party-affairs/pkg/util"
)

type User struct {
	conf *config.Config
}

func NewUser(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

func (l *User) Page(ctx kratosx.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	user := model.User{}
	list, total, err := user.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Name != nil {
				db = db.Where("name=?", *in.Name)
			}
			if in.Phone != nil {
				db = db.Where("phone=?", *in.Phone)
			}
			if in.Email != nil {
				db = db.Where("email=?", *in.Email)
			}
			if in.Status != nil {
				db = db.Where("status=?", *in.Status)
			}
			if in.StartTime != nil {
				db = db.Where("created_at>?", *in.StartTime)
			}
			if in.EndTime != nil {
				db = db.Where("created_at<?", *in.EndTime)
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageUserReply{Total: total}
	if err := util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Current 获取当前用户信息
func (l *User) Current(ctx kratosx.Context) (*v1.User, error) {
	user := model.User{}
	if err := user.OneByID(ctx, md.UserId(ctx)); err != nil {
		return nil, v1.NotFoundError()
	}

	reply := v1.User{}
	if err := util.Transform(user, &reply); err != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

// ChangeStatus 改变用户状态
func (l *User) ChangeStatus(ctx kratosx.Context, in *v1.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	user := model.User{
		BaseModel: model.BaseModel{ID: in.Id},
		Status:    &in.Status,
	}

	if err := user.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

func (l *User) Add(ctx kratosx.Context, in *v1.AddUserRequest) (*empty.Empty, error) {
	user := model.User{
		Nickname: in.Name,
		Avatar:   l.conf.DefaultUserAvatar,
		Status:   proto.Bool(true),
	}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Create(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *User) Update(ctx kratosx.Context, in *v1.UpdateUserRequest) (*empty.Empty, error) {
	user := model.User{}
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	if err := user.Update(ctx); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}

	return nil, nil
}

func (l *User) Delete(ctx kratosx.Context, in *v1.DeleteUserRequest) (*empty.Empty, error) {
	user := model.User{}
	if err := user.DeleteByID(ctx, in.Id); err != nil {
		return nil, v1.DataErrorFormat(err.Error())
	}
	return nil, nil
}
