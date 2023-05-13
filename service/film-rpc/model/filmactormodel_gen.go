// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	filmActorFieldNames          = builder.RawFieldNames(&FilmActor{})
	filmActorRows                = strings.Join(filmActorFieldNames, ",")
	filmActorRowsExpectAutoSet   = strings.Join(stringx.Remove(filmActorFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	filmActorRowsWithPlaceHolder = strings.Join(stringx.Remove(filmActorFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheFilmActorIdPrefix = "cache:filmActor:id:"
)

type (
	filmActorModel interface {
		Insert(ctx context.Context, data *FilmActor) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FilmActor, error)
		Update(ctx context.Context, data *FilmActor) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFilmActorModel struct {
		sqlc.CachedConn
		table string
	}

	FilmActor struct {
		Id       int64  `db:"id"`        // 主键编号
		FilmId   int64  `db:"film_id"`   // 影片编号,对应mooc_film_t
		ActorId  int64  `db:"actor_id"`  // 演员编号,对应mooc_actor_t
		RoleName string `db:"role_name"` // 角色名称
	}
)

func newFilmActorModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFilmActorModel {
	return &defaultFilmActorModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`film_actor`",
	}
}

func (m *defaultFilmActorModel) Delete(ctx context.Context, id int64) error {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, filmActorIdKey)
	return err
}

func (m *defaultFilmActorModel) FindOne(ctx context.Context, id int64) (*FilmActor, error) {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, id)
	var resp FilmActor
	err := m.QueryRowCtx(ctx, &resp, filmActorIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", filmActorRows, m.table)
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

func (m *defaultFilmActorModel) Insert(ctx context.Context, data *FilmActor) (sql.Result, error) {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, filmActorRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.FilmId, data.ActorId, data.RoleName)
	}, filmActorIdKey)
	return ret, err
}

func (m *defaultFilmActorModel) Update(ctx context.Context, data *FilmActor) error {
	filmActorIdKey := fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, filmActorRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FilmId, data.ActorId, data.RoleName, data.Id)
	}, filmActorIdKey)
	return err
}

func (m *defaultFilmActorModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFilmActorIdPrefix, primary)
}

func (m *defaultFilmActorModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", filmActorRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFilmActorModel) tableName() string {
	return m.table
}
