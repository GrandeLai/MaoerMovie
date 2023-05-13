package model

import (
	"MaoerMovie/service/user-rpc/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoryModel = (*customCategoryModel)(nil)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
		FindAll(ctx context.Context) ([]*Category, error)
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn, c),
	}
}

func (m *defaultCategoryModel) FindAll(ctx context.Context) ([]*Category, error) {
	var resp []*Category
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCategoryModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(categoryRows).From(m.table)
}
