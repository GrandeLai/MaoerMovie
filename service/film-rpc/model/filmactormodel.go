package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var _ FilmActorModel = (*customFilmActorModel)(nil)

var filmActorRowsExpectAutoSetButId = strings.Join(stringx.Remove(filmActorFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// FilmActorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmActorModel.
	FilmActorModel interface {
		filmActorModel
		FindAllByFilmId(ctx context.Context, filmId int64) ([]*FilmActor, error)
		FindOneByFilmId(ctx context.Context, filmId int64) (*FilmActor, error)
		InsertWithNewId(ctx context.Context, data *FilmActor) (sql.Result, error)
		DeleteByFilmId(ctx context.Context, filmId int64) error
	}

	customFilmActorModel struct {
		*defaultFilmActorModel
	}
)

// NewFilmActorModel returns a model for the database table.
func NewFilmActorModel(conn sqlx.SqlConn, c cache.CacheConf) FilmActorModel {
	return &customFilmActorModel{
		defaultFilmActorModel: newFilmActorModel(conn, c),
	}
}

func (m *defaultFilmActorModel) FindAllByFilmId(ctx context.Context, filmId int64) ([]*FilmActor, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("film_id = ?", filmId).ToSql()
	var resp []*FilmActor
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFilmActorModel) FindOneByFilmId(ctx context.Context, filmId int64) (*FilmActor, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("film_id = ?", filmId).Limit(1).ToSql()
	var resp FilmActor
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

func (m *defaultFilmActorModel) InsertWithNewId(ctx context.Context, data *FilmActor) (sql.Result, error) {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, filmActorRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.FilmId, data.ActorId, data.RoleName)
	}, filmActorIdKey)
	return ret, err
}

func (m *defaultFilmActorModel) DeleteByFilmId(ctx context.Context, filmId int64) error {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, filmId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `film_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, filmId)
	}, filmActorIdKey)
	return err
}

func (m *defaultFilmActorModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(filmActorRows).From(m.table)
}
