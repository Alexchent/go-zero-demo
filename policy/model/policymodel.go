package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PolicyModel = (*customPolicyModel)(nil)

type (
	// PolicyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPolicyModel.
	PolicyModel interface {
		policyModel
		withSession(session sqlx.Session) PolicyModel
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
