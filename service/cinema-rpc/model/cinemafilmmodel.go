package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CinemaFilmModel = (*customCinemaFilmModel)(nil)

type (
	// CinemaFilmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCinemaFilmModel.
	CinemaFilmModel interface {
		cinemaFilmModel
		FindAllByCinemaId(ctx context.Context, cinemaId int64) ([]*CinemaFilm, error)
	}

	customCinemaFilmModel struct {
		*defaultCinemaFilmModel
	}
)

// NewCinemaFilmModel returns a model for the database table.
func NewCinemaFilmModel(conn sqlx.SqlConn, c cache.CacheConf) CinemaFilmModel {
	return &customCinemaFilmModel{
		defaultCinemaFilmModel: newCinemaFilmModel(conn, c),
	}
}

func (m *defaultCinemaFilmModel) FindAllByCinemaId(ctx context.Context, cinemaId int64) ([]*CinemaFilm, error) {
	var resp []*CinemaFilm
	rowBuilder := m.RowBuilder().Where("cinema_id=?", cinemaId)
	query, values, err := rowBuilder.ToSql()
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

func (m *defaultCinemaFilmModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(cinemaFilmRows).From(m.table)
}
