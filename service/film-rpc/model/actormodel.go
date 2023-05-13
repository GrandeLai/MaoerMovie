package model

import (
	"MaoerMovie/service/user-rpc/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActorModel = (*customActorModel)(nil)

type (
	// ActorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActorModel.
	ActorModel interface {
		actorModel
		FindAllInPage(ctx context.Context, startIndex int64, pageSize int64) ([]*Actor, error)
		FindByActorName(ctx context.Context, name string) (*Actor, error)
	}

	customActorModel struct {
		*defaultActorModel
	}
)

// NewActorModel returns a model for the database table.
func NewActorModel(conn sqlx.SqlConn, c cache.CacheConf) ActorModel {
	return &customActorModel{
		defaultActorModel: newActorModel(conn, c),
	}
}

func (m *defaultActorModel) FindAllInPage(ctx context.Context, startIndex int64, pageSize int64) ([]*Actor, error) {
	var resp []*Actor
	rowBuilder := m.RowBuilder()
	if startIndex != 0 || pageSize != 0 {
		rowBuilder = rowBuilder.Offset(uint64(startIndex - 1)).Limit(uint64(pageSize))
	}
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

func (m *defaultActorModel) FindByActorName(ctx context.Context, name string) (*Actor, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("actor_name = ?", name).Limit(1).ToSql()
	var resp Actor
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultActorModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(actorRows).From(m.table)
}

func (m *defaultActorModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
