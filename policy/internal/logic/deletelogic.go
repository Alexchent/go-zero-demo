package logic

import (
	"context"
	"database/sql"
	"errors"
	"go-zero-demo/policy/model"
	"time"

	"go-zero-demo/policy/internal/svc"
	pb "go-zero-demo/policy/pb/policy"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *pb.DeleteReq) (*pb.Placeholder, error) {
	one, err := l.svcCtx.PolicyModel.FindOne(l.ctx, uint64(in.GetID()))
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	one.Updated = time.Now()
	one.Deleted = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = l.svcCtx.PolicyModel.Update(l.ctx, one)
	return &pb.Placeholder{}, err
}
