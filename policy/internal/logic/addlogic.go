package logic

import (
	"context"
	"database/sql"
	"encoding/json"
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
	rules, err := json.Marshal(in.GetRule())
	if err != nil {
		return nil, err
	}
	logx.Info("rules:", rules)
	insert, err := l.svcCtx.PolicyModel.Insert(l.ctx, &model.Policy{
		Cate: in.GetCate(),
		Attr: in.GetAttr(),
		Rule: sql.NullString{
			String: string(rules),
			Valid:  true,
		},
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
