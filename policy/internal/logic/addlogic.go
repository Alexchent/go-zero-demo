package logic

import (
	"context"
	"go-zero-demo/policy/internal/svc"
	"go-zero-demo/policy/model"
	"go-zero-demo/policy/pb/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.Input) (*pb.AddRep, error) {
	insert, err := l.svcCtx.PolicyModel.Insert(l.ctx, &model.Policy{
		Cate:    in.GetCate(),
		Attr:    in.GetAttr(),
		Rule:    in.GetRule(),
		Created: time.Now(),
		Updated: time.Now(),
		State:   0,
	})
	if err != nil {
		return nil, err
	}
	id, err := insert.LastInsertId()
	return &pb.AddRep{ID: id}, err
}
