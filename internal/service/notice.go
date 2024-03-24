package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	userV1 "github.com/limes-cloud/user-center/api/user/v1"
	"google.golang.org/protobuf/proto"

	v1 "party-affairs/api/v1"
	"party-affairs/internal/biz"
	"party-affairs/internal/biz/types"
	"party-affairs/internal/consts"
	"party-affairs/internal/pkg/service"
)

func (s *Service) PageNotice(ctx context.Context, in *v1.PageNoticeRequest) (*v1.PageNoticeReply, error) {
	list, total, err := s.notice.Page(kratosx.MustContext(ctx), &types.PageNoticeRequest{
		Page:     in.Page,
		PageSize: in.PageSize,
		Title:    in.Title,
		NotRead:  in.NotRead,
	})
	if err != nil {
		return nil, err
	}

	reply := v1.PageNoticeReply{Total: total}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	for ind, item := range list {
		reply.List[ind].IsRead = item.NoticeUser != nil
	}

	return &reply, nil
}

func (s *Service) PageNoticeUser(ctx context.Context, in *v1.PageNoticeUserRequest) (*v1.PageNoticeUserReply, error) {
	ids, err := s.notice.ReadUserIds(kratosx.MustContext(ctx), in.NoticeId)
	if err != nil {
		return nil, err
	}

	client, err := service.NewUser(ctx)
	if err != nil {
		return nil, v1.UserCenterError()
	}
	req := &userV1.PageUserRequest{
		App:      proto.String(consts.ClientApp),
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	if in.IsRead {
		req.InIds = ids
	} else {
		req.NotInIds = ids
	}
	resp, err := client.PageUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.PageNoticeUserReply{
		Total: resp.Total,
		List:  resp.List,
	}, nil
}

func (s *Service) GetNotice(ctx context.Context, in *v1.GetNoticeRequest) (*v1.Notice, error) {
	nc, err := s.notice.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}

	reply := v1.Notice{}
	if err := copier.Copy(&reply, nc); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) ReadNotice(ctx context.Context, in *v1.ReadNoticeRequest) (*empty.Empty, error) {
	return nil, s.notice.Read(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) AddNotice(ctx context.Context, in *v1.AddNoticeRequest) (*empty.Empty, error) {
	var notice biz.Notice
	kCtx := kratosx.MustContext(ctx)
	if err := copier.Copy(&notice, in); err != nil {
		return nil, v1.TransformError()
	}

	if _, err := s.notice.Add(kCtx, &notice); err != nil {
		return nil, err
	}
	go func() {
		if err := s.notice.SendNoticeEmail(kratosx.MustContext(context.Background()), notice.ID); err != nil {
			kCtx.Logger().Errorf("发送邮件失败：%s", err.Error())
		}
	}()
	return nil, nil
}

func (s *Service) UpdateNotice(ctx context.Context, in *v1.UpdateNoticeRequest) (*empty.Empty, error) {
	var notice biz.Notice
	if err := copier.Copy(&notice, in); err != nil {
		return nil, v1.TransformError()
	}

	kCtx := kratosx.MustContext(ctx)
	if err := s.notice.Update(kCtx, &notice); err != nil {
		return nil, nil
	}
	go func() {
		if err := s.notice.SendNoticeEmail(kratosx.MustContext(context.Background()), notice.ID); err != nil {
			kCtx.Logger().Errorf("发送邮件失败：%s", err.Error())
		}
	}()
	return nil, s.notice.Update(kratosx.MustContext(ctx), &notice)
}

func (s *Service) DeleteNotice(ctx context.Context, in *v1.DeleteNoticeRequest) (*empty.Empty, error) {
	return nil, s.notice.Delete(kratosx.MustContext(ctx), in.Id)
}
