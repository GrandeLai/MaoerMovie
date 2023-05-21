package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PayModel = (*customPayModel)(nil)

type (
	// PayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayModel.
	PayModel interface {
		payModel
		TxUpdate(tx *sql.Tx, data *Pay) (sql.Result, error)
		FindByPaySnAndUserId(ctx context.Context, paySn string, userId int64) (*Pay, error)
		FindByPaySn(ctx context.Context, paySn string) (*Pay, error)
	}

	customPayModel struct {
		*defaultPayModel
	}
)

// NewPayModel returns a model for the database table.
func NewPayModel(conn sqlx.SqlConn, c cache.CacheConf) PayModel {
	return &customPayModel{
		defaultPayModel: newPayModel(conn, c),
	}
}

func (m *defaultPayModel) TxUpdate(tx *sql.Tx, data *Pay) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, payRowsWithPlaceHolder)
	ret, err := tx.Exec(query, data.PaySn, data.UserId, data.OrderId, data.BuyerAccount, data.Price, data.Subject, data.Status, data.Id)
	return ret, err
}

func (m *defaultPayModel) FindByPaySnAndUserId(ctx context.Context, paySn string, userId int64) (*Pay, error) {
	commentIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, paySn)
	var resp Pay
	err := m.QueryRowCtx(ctx, &resp, commentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `pay_sn` = ? AND `user_id = ?`  limit 1", payRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, paySn, userId)
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

func (m *defaultPayModel) FindByPaySn(ctx context.Context, paySn string) (*Pay, error) {
	commentIdKey := fmt.Sprintf("%s%v", cachePayIdPrefix, paySn)
	var resp Pay
	err := m.QueryRowCtx(ctx, &resp, commentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `pay_sn` = ? limit 1", payRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, paySn)
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
