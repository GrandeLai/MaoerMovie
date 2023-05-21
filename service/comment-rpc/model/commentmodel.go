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

var _ CommentModel = (*customCommentModel)(nil)

var commentRowsExpectAutoSetButId = strings.Join(stringx.Remove(commentFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindAllInPageByFilmId(ctx context.Context, filmId int64, startIndex int64, pageSize int64) ([]*Comment, error)
		InsertWithNewId(ctx context.Context, data *Comment) (sql.Result, error)
		FindOneByUserIdAndCommentId(ctx context.Context, userId int64, commentId int64) (*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

func (m *defaultCommentModel) FindAllInPageByFilmId(ctx context.Context, filmId int64, startIndex int64, pageSize int64) ([]*Comment, error) {
	var resp []*Comment
	rowBuilder := m.RowBuilder().Where("film_id = ?", filmId)
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

func (m *defaultCommentModel) InsertWithNewId(ctx context.Context, data *Comment) (sql.Result, error) {
	filmIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.FilmId, data.Content, data.Score)
	}, filmIdKey)
	return ret, err
}

func (m *defaultCommentModel) FindOneByUserIdAndCommentId(ctx context.Context, userId int64, commentId int64) (*Comment, error) {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, commentId)
	var resp Comment
	err := m.QueryRowCtx(ctx, &resp, commentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? AND `user_id` = ? limit 1", commentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, commentId, userId)
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

func (m *defaultCommentModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(commentRows).From(m.table)
}

func (m *defaultCommentModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
