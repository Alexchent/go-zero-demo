package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PolicyModel = (*customPolicyModel)(nil)

type (
	// PolicyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPolicyModel.
	PolicyModel interface {
		policyModel
		withSession(session sqlx.Session) PolicyModel
		FindByCate(ctx context.Context, cate string) ([]Policy, error)
	}

	customPolicyModel struct {
		*defaultPolicyModel
	}
)

// NewPolicyModel returns a model for the database table.
func NewPolicyModel(conn sqlx.SqlConn) PolicyModel {
	return &customPolicyModel{
		defaultPolicyModel: newPolicyModel(conn),
	}
}

func (m *customPolicyModel) withSession(session sqlx.Session) PolicyModel {
	return NewPolicyModel(sqlx.NewSqlConnFromSession(session))
}

func (m *defaultPolicyModel) FindByCate(ctx context.Context, cate string) ([]Policy, error) {
	query := fmt.Sprintf("select %s from %s where `cate` = ? ", policyRows, m.table)
	resp := make([]Policy, 0)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, cate)
	return resp, err
}
