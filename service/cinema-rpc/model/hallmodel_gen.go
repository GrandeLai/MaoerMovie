// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	hallFieldNames          = builder.RawFieldNames(&Hall{})
	hallRows                = strings.Join(hallFieldNames, ",")
	hallRowsExpectAutoSet   = strings.Join(stringx.Remove(hallFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	hallRowsWithPlaceHolder = strings.Join(stringx.Remove(hallFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"

	cacheHallIdPrefix = "cache:hall:id:"
)

type (
	hallModel interface {
		Insert(ctx context.Context, data *Hall) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Hall, error)
		Update(ctx context.Context, data *Hall) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHallModel struct {
		sqlc.CachedConn
		table string
	}

	Hall struct {
		Id          int64  `db:"id"`           // 主键编号
		HallName    string `db:"hall_name"`    // 显示名称
		SeatAddress string `db:"seat_address"` // 座位文件存放地址
	}
)

func newHallModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHallModel {
	return &defaultHallModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`hall`",
	}
}

func (m *defaultHallModel) Delete(ctx context.Context, id int64) error {
	hallIdKey := fmt.Sprintf("%s%v", cacheHallIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, hallIdKey)
	return err
}

func (m *defaultHallModel) FindOne(ctx context.Context, id int64) (*Hall, error) {
	hallIdKey := fmt.Sprintf("%s%v", cacheHallIdPrefix, id)
	var resp Hall
	err := m.QueryRowCtx(ctx, &resp, hallIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", hallRows, m.table)
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

func (m *defaultHallModel) Insert(ctx context.Context, data *Hall) (sql.Result, error) {
	hallIdKey := fmt.Sprintf("%s%v", cacheHallIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, hallRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.HallName, data.SeatAddress)
	}, hallIdKey)
	return ret, err
}

func (m *defaultHallModel) Update(ctx context.Context, data *Hall) error {
	hallIdKey := fmt.Sprintf("%s%v", cacheHallIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, hallRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.HallName, data.SeatAddress, data.Id)
	}, hallIdKey)
	return err
}

func (m *defaultHallModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheHallIdPrefix, primary)
}

func (m *defaultHallModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", hallRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHallModel) tableName() string {
	return m.table
}
