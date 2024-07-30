// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	policyFieldNames          = builder.RawFieldNames(&Policy{})
	policyRows                = strings.Join(policyFieldNames, ",")
	policyRowsExpectAutoSet   = strings.Join(stringx.Remove(policyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	policyRowsWithPlaceHolder = strings.Join(stringx.Remove(policyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	policyModel interface {
		Insert(ctx context.Context, data *Policy) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Policy, error)
		Update(ctx context.Context, data *Policy) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultPolicyModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Policy struct {
		Id      uint64         `db:"id"`
		Cate    string         `db:"cate"`  // 类别
		Attr    sql.NullString `db:"attr"`  // 对象信息
		Rule    sql.NullString `db:"rule"`  // 匹配规则
		State   int64          `db:"state"` // 状态0关闭 1开启
		Created sql.NullTime   `db:"created"`
		Updated sql.NullTime   `db:"updated"`
		Deleted sql.NullTime   `db:"deleted"`
	}
)

func newPolicyModel(conn sqlx.SqlConn) *defaultPolicyModel {
	return &defaultPolicyModel{
		conn:  conn,
		table: "`policy`",
	}
}

func (m *defaultPolicyModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPolicyModel) FindOne(ctx context.Context, id uint64) (*Policy, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", policyRows, m.table)
	var resp Policy
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPolicyModel) Insert(ctx context.Context, data *Policy) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, policyRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Cate, data.Attr, data.Rule, data.State, data.Created, data.Updated, data.Deleted)
	return ret, err
}

func (m *defaultPolicyModel) Update(ctx context.Context, data *Policy) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, policyRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Cate, data.Attr, data.Rule, data.State, data.Created, data.Updated, data.Deleted, data.Id)
	return err
}

func (m *defaultPolicyModel) tableName() string {
	return m.table
}
