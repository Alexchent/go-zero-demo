// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	policyFieldNames          = builder.RawFieldNames(&Policy{})
	policyRows                = strings.Join(policyFieldNames, ",")
	policyRowsExpectAutoSet   = strings.Join(stringx.Remove(policyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	policyRowsWithPlaceHolder = strings.Join(stringx.Remove(policyFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cachePolicyIdPrefix = "cache:policy:id:"
)

type (
	policyModel interface {
		Insert(ctx context.Context, data *Policy) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Policy, error)
		Update(ctx context.Context, data *Policy) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultPolicyModel struct {
		sqlc.CachedConn
		table string
	}

	Policy struct {
		Id      uint64       `db:"id"`
		Cate    string       `db:"cate"`  // 类别
		Attr    string       `db:"attr"`  // 对象信息
		Rule    string       `db:"rule"`  // 匹配规则
		State   int64        `db:"state"` // 状态0关闭 1开启
		Created time.Time    `db:"created"`
		Updated time.Time    `db:"updated"`
		Deleted sql.NullTime `db:"deleted"`
	}
)

func newPolicyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPolicyModel {
	return &defaultPolicyModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`policy`",
	}
}

func (m *defaultPolicyModel) Delete(ctx context.Context, id uint64) error {
	policyIdKey := fmt.Sprintf("%s%v", cachePolicyIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, policyIdKey)
	return err
}

func (m *defaultPolicyModel) FindOne(ctx context.Context, id uint64) (*Policy, error) {
	policyIdKey := fmt.Sprintf("%s%v", cachePolicyIdPrefix, id)
	var resp Policy
	err := m.QueryRowCtx(ctx, &resp, policyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", policyRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPolicyModel) Insert(ctx context.Context, data *Policy) (sql.Result, error) {
	policyIdKey := fmt.Sprintf("%s%v", cachePolicyIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, policyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Cate, data.Attr, data.Rule, data.State, data.Created, data.Updated, data.Deleted)
	}, policyIdKey)
	return ret, err
}

func (m *defaultPolicyModel) Update(ctx context.Context, data *Policy) error {
	policyIdKey := fmt.Sprintf("%s%v", cachePolicyIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, policyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Cate, data.Attr, data.Rule, data.State, data.Created, data.Updated, data.Deleted, data.Id)
	}, policyIdKey)
	return err
}

func (m *defaultPolicyModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePolicyIdPrefix, primary)
}

func (m *defaultPolicyModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", policyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPolicyModel) tableName() string {
	return m.table
}
