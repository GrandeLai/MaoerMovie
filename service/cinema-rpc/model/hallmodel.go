package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HallModel = (*customHallModel)(nil)

type (
	// HallModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHallModel.
	HallModel interface {
		hallModel
		FindAll(ctx context.Context) ([]*Hall, error)
	}

	customHallModel struct {
		*defaultHallModel
	}
)

// NewHallModel returns a model for the database table.
func NewHallModel(conn sqlx.SqlConn, c cache.CacheConf) HallModel {
	return &customHallModel{
		defaultHallModel: newHallModel(conn, c),
	}
}

func (m *defaultHallModel) FindAll(ctx context.Context) ([]*Hall, error) {
	var resp []*Hall
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

func (m *defaultHallModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(hallRows).From(m.table)
}
