// Code generated by goctl. DO NOT EDIT!

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
	showFieldNames          = builder.RawFieldNames(&Show{})
	showRows                = strings.Join(showFieldNames, ",")
	showRowsExpectAutoSet   = strings.Join(stringx.Remove(showFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	showRowsWithPlaceHolder = strings.Join(stringx.Remove(showFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheShowIdPrefix = "cache:show:id:"
)

type (
	showModel interface {
		Insert(ctx context.Context, data *Show) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Show, error)
		Update(ctx context.Context, data *Show) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShowModel struct {
		sqlc.CachedConn
		table string
	}

	Show struct {
		Id           int64     `db:"id"`            // 主键编号
		CinemaId     int64     `db:"cinema_id"`     // 影院编号
		FilmId       int64     `db:"film_id"`       // 电影编号
		BeginTime    string    `db:"begin_time"`    // 开始时间
		EndTime      string    `db:"end_time"`      // 结束时间
		HallId       int64     `db:"hall_id"`       // 放映厅类型编号
		Price        float64   `db:"price"`         // 票价
		Date         time.Time `db:"date"`          // 放映日期
		FilmLanguage string    `db:"film_language"` // 电影语言
		SoldNum      int64     `db:"sold_num"`      // 放映厅类型编号
	}
)

func newShowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultShowModel {
	return &defaultShowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`show`",
	}
}

func (m *defaultShowModel) Delete(ctx context.Context, id int64) error {
	showIdKey := fmt.Sprintf("%s%v", cacheShowIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, showIdKey)
	return err
}

func (m *defaultShowModel) FindOne(ctx context.Context, id int64) (*Show, error) {
	showIdKey := fmt.Sprintf("%s%v", cacheShowIdPrefix, id)
	var resp Show
	err := m.QueryRowCtx(ctx, &resp, showIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", showRows, m.table)
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

func (m *defaultShowModel) Insert(ctx context.Context, data *Show) (sql.Result, error) {
	showIdKey := fmt.Sprintf("%s%v", cacheShowIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, showRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CinemaId, data.FilmId, data.BeginTime, data.EndTime, data.HallId, data.Price, data.Date, data.FilmLanguage, data.SoldNum)
	}, showIdKey)
	return ret, err
}

func (m *defaultShowModel) Update(ctx context.Context, data *Show) error {
	showIdKey := fmt.Sprintf("%s%v", cacheShowIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, showRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CinemaId, data.FilmId, data.BeginTime, data.EndTime, data.HallId, data.Price, data.Date, data.FilmLanguage, data.SoldNum, data.Id)
	}, showIdKey)
	return err
}

func (m *defaultShowModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheShowIdPrefix, primary)
}

func (m *defaultShowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", showRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShowModel) tableName() string {
	return m.table
}
