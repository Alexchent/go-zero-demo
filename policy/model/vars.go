package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	Policy_STATE_ON  = 1
	Policy_STATE_OFF = 0
)
