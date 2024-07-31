package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PolicyModel = (*customPolicyModel)(nil)

type (
	// PolicyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPolicyModel.
	PolicyModel interface {
		policyModel
		FindByCate(ctx context.Context, cate string) ([]Policy, error)
	}

	customPolicyModel struct {
		*defaultPolicyModel
	}
)

// NewPolicyModel returns a model for the database table.
func NewPolicyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PolicyModel {
	return &customPolicyModel{
		defaultPolicyModel: newPolicyModel(conn, c, opts...),
	}
}

func (m *customPolicyModel) FindByCate(ctx context.Context, cate string) ([]Policy, error) {
	query := fmt.Sprintf("select %s from %s where `state` = ? AND `cate` = ? ", policyRows, m.table)
	resp := make([]Policy, 0)
	err := m.QueryRowsNoCache(&resp, query, Policy_STATE_ON, cate)
	return resp, err
}
