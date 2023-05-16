package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BrandModel = (*customBrandModel)(nil)

type (
	// BrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBrandModel.
	BrandModel interface {
		brandModel
		FindAll(ctx context.Context) ([]*Brand, error)
	}

	customBrandModel struct {
		*defaultBrandModel
	}
)

// NewBrandModel returns a model for the database table.
func NewBrandModel(conn sqlx.SqlConn, c cache.CacheConf) BrandModel {
	return &customBrandModel{
		defaultBrandModel: newBrandModel(conn, c),
	}
}

func (m *defaultBrandModel) FindAll(ctx context.Context) ([]*Brand, error) {
	var resp []*Brand
	query, values, err := m.RowBuilder().ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBrandModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(brandRows).From(m.table)
}
