// Code generated by goctl. DO NOT EDIT.
// Source: policy.proto

package server

import (
	"context"

	"go-zero-demo/policy/internal/logic"
	"go-zero-demo/policy/internal/svc"
	"go-zero-demo/policy/pb/pb"
)

type PolicyServer struct {
	svcCtx *svc.ServiceContext
	policy.UnimplementedPolicyServer
}

func NewPolicyServer(svcCtx *svc.ServiceContext) *PolicyServer {
	return &PolicyServer{
		svcCtx: svcCtx,
	}
}

func (s *PolicyServer) Add(ctx context.Context, in *policy.Input) (*policy.AddRep, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *PolicyServer) Update(ctx context.Context, in *policy.Input) (*policy.AddRep, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *PolicyServer) Delete(ctx context.Context, in *policy.DeleteReq) (*policy.Placeholder, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}

func (s *PolicyServer) Check(ctx context.Context, in *policy.CheckReq) (*policy.CheckRep, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
