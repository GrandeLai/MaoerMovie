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

var _ FilmInfoModel = (*customFilmInfoModel)(nil)

var filmInfoRowsExpectAutoSetButId = strings.Join(stringx.Remove(filmFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// FilmInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmInfoModel.
	FilmInfoModel interface {
		filmInfoModel
		InsertWithNewId(ctx context.Context, data *FilmInfo) (sql.Result, error)
		FindOneByFilmId(ctx context.Context, id int64) (*FilmInfo, error)
		DeleteByFilmId(ctx context.Context, id int64) error
	}

	customFilmInfoModel struct {
		*defaultFilmInfoModel
	}
)

// NewFilmInfoModel returns a model for the database table.
func NewFilmInfoModel(conn sqlx.SqlConn, c cache.CacheConf) FilmInfoModel {
	return &customFilmInfoModel{
		defaultFilmInfoModel: newFilmInfoModel(conn, c),
	}
}

func (m *defaultFilmInfoModel) InsertWithNewId(ctx context.Context, data *FilmInfo) (sql.Result, error) {
	filmInfoIdKey := fmt.Sprintf("%s%v", cacheFilmInfoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, filmInfoRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.FilmId, data.FilmPreSaleNum, data.FilmBoxOffice, data.FilmImgs)
	}, filmInfoIdKey)
	return ret, err
}

func (m *defaultFilmInfoModel) FindOneByFilmId(ctx context.Context, id int64) (*FilmInfo, error) {
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("film_id = ?", id).Limit(1).ToSql()
	var resp FilmInfo
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

func (m *defaultFilmInfoModel) DeleteByFilmId(ctx context.Context, id int64) error {
	filmInfoIdKey := fmt.Sprintf("%s%v", cacheFilmInfoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `film_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, filmInfoIdKey)
	return err
}

func (m *defaultFilmInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(filmInfoRows).From(m.table)
}

func (m *defaultFilmInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
