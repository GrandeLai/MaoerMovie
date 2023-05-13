package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var _ FilmModel = (*customFilmModel)(nil)

var filmRowsExpectAutoSetButId = strings.Join(stringx.Remove(filmFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// FilmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFilmModel.
	FilmModel interface {
		filmModel
		FindAllInPage(ctx context.Context, startIndex int64, pageSize int64) ([]*Film, error)
		InsertWithNewId(ctx context.Context, data *Film) (sql.Result, error)
	}

	customFilmModel struct {
		*defaultFilmModel
	}
)

// NewFilmModel returns a model for the database table.
func NewFilmModel(conn sqlx.SqlConn, c cache.CacheConf) FilmModel {
	return &customFilmModel{
		defaultFilmModel: newFilmModel(conn, c),
	}
}

func (m *defaultFilmModel) FindAllInPage(ctx context.Context, startIndex int64, pageSize int64) ([]*Film, error) {
	var resp []*Film
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
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFilmModel) InsertWithNewId(ctx context.Context, data *Film) (sql.Result, error) {
	filmIdKey := fmt.Sprintf("%s%v", cacheFilmIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)", m.table, filmRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.FilmName, data.FilmEnName, data.FilmType, data.FilmCover, data.FilmLength, data.CategoryId, data.FilmArea, data.FilmTime, data.DirectorId, data.Biography)
	}, filmIdKey)
	return ret, err
}

func (m *defaultFilmModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(filmRows).From(m.table)
}

func (m *defaultFilmModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
