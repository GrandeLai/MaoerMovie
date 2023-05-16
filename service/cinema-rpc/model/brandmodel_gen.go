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
	brandFieldNames          = builder.RawFieldNames(&Brand{})
	brandRows                = strings.Join(brandFieldNames, ",")
	brandRowsExpectAutoSet   = strings.Join(stringx.Remove(brandFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	brandRowsWithPlaceHolder = strings.Join(stringx.Remove(brandFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"

	cacheBrandIdPrefix = "cache:brand:id:"
)

type (
	brandModel interface {
		Insert(ctx context.Context, data *Brand) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Brand, error)
		Update(ctx context.Context, data *Brand) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBrandModel struct {
		sqlc.CachedConn
		table string
	}

	Brand struct {
		Id        int64  `db:"id"`         // 主键编号
		BrandName string `db:"brand_name"` // 品牌名称
	}
)

func newBrandModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBrandModel {
	return &defaultBrandModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`brand`",
	}
}

func (m *defaultBrandModel) Delete(ctx context.Context, id int64) error {
	brandIdKey := fmt.Sprintf("%s%v", cacheBrandIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, brandIdKey)
	return err
}

func (m *defaultBrandModel) FindOne(ctx context.Context, id int64) (*Brand, error) {
	brandIdKey := fmt.Sprintf("%s%v", cacheBrandIdPrefix, id)
	var resp Brand
	err := m.QueryRowCtx(ctx, &resp, brandIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", brandRows, m.table)
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

func (m *defaultBrandModel) Insert(ctx context.Context, data *Brand) (sql.Result, error) {
	brandIdKey := fmt.Sprintf("%s%v", cacheBrandIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, brandRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.BrandName)
	}, brandIdKey)
	return ret, err
}

func (m *defaultBrandModel) Update(ctx context.Context, data *Brand) error {
	brandIdKey := fmt.Sprintf("%s%v", cacheBrandIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, brandRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.BrandName, data.Id)
	}, brandIdKey)
	return err
}

func (m *defaultBrandModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBrandIdPrefix, primary)
}

func (m *defaultBrandModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", brandRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBrandModel) tableName() string {
	return m.table
}