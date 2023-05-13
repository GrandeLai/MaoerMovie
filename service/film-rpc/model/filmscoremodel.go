package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FilmScoreModel = (*customFilmScoreModel)(nil)

type (
	// FilmScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmScoreModel.
	FilmScoreModel interface {
		filmScoreModel
		FindOneByFilmId(ctx context.Context, id int64) (*FilmScore, error)
		DeleteByFilmId(ctx context.Context, id int64) error
	}

	customFilmScoreModel struct {
		*defaultFilmScoreModel
	}
)

// NewFilmScoreModel returns a model for the database table.
func NewFilmScoreModel(conn sqlx.SqlConn, c cache.CacheConf) FilmScoreModel {
	return &customFilmScoreModel{
		defaultFilmScoreModel: newFilmScoreModel(conn, c),
	}
}

func (m *defaultFilmScoreModel) FindOneByFilmId(ctx context.Context, id int64) (*FilmScore, error) {
	filmScoreFilmIdKey := fmt.Sprintf("%s%v", cacheFilmScoreFilmIdPrefix, id)
	var resp FilmScore
	err := m.QueryRowCtx(ctx, &resp, filmScoreFilmIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `film_id` = ? limit 1", filmScoreRows, m.table)
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

func (m *defaultFilmScoreModel) DeleteByFilmId(ctx context.Context, id int64) error {
	filmScoreIdKey := fmt.Sprintf("%s%v", cacheFilmScoreIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `film_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, filmScoreIdKey)
	return err
}
