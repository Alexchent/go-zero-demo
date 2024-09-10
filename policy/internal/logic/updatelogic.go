package logic

import (
	"context"
	"go-zero-demo/policy/internal/svc"
	"go-zero-demo/policy/model"
	"go-zero-demo/policy/pb/policy"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *policy.Input) (*policy.UpdateResp, error) {

	err := l.svcCtx.PolicyModel.Update(l.ctx, &model.Policy{
		Id:      uint64(in.GetID()),
		Cate:    in.GetCate(),
		Attr:    in.GetAttr(),
		Rule:    in.GetRule(),
		State:   in.GetState(),
		Updated: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &policy.UpdateResp{}, err
}
