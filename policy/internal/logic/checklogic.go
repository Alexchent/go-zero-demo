package logic

import (
	"context"
	"github.com/expr-lang/expr"

	"go-zero-demo/policy/internal/svc"
	"go-zero-demo/policy/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *pb.CheckReq) (*pb.CheckRep, error) {
	policy, err := l.svcCtx.PolicyModel.FindByCate(l.ctx, in.GetCate())
	if err != nil {
		return nil, err
	}
	data := make([]*pb.Input, 0)
	for _, p := range policy {
		logx.Info("expr:", p.Rule, in.GetUser())
		program, _ := expr.Compile(p.Rule, expr.AsBool())
		out, err := expr.Run(program, in.GetUser())
		if err != nil {
			logx.Error("expr-err:" + err.Error())
		}
		if out == true {
			data = append(data, &pb.Input{
				Cate: p.Cate,
				Attr: p.Attr,
				Rule: p.Rule,
			})
		}
	}

	return &pb.CheckRep{List: data}, nil
}
