// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	actorFieldNames          = builder.RawFieldNames(&Actor{})
	actorRows                = strings.Join(actorFieldNames, ",")
	actorRowsExpectAutoSet   = strings.Join(stringx.Remove(actorFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	actorRowsWithPlaceHolder = strings.Join(stringx.Remove(actorFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cacheActorIdPrefix = "cache:actor:id:"
)

type (
	actorModel interface {
		Insert(ctx context.Context, data *Actor) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Actor, error)
		Update(ctx context.Context, data *Actor) error
		Delete(ctx context.Context, id int64) error
	}

	defaultActorModel struct {
		sqlc.CachedConn
		table string
	}

	Actor struct {
		Id         int64     `db:"id"`         // 主键编号
		ActorName  string    `db:"actor_name"` // 演员名称
		ActorImg   string    `db:"actor_img"`  // 演员图片位置
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newActorModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActorModel {
	return &defaultActorModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`actor`",
	}
}

func (m *defaultActorModel) Delete(ctx context.Context, id int64) error {
	actorIdKey := fmt.Sprintf("%s%v", cacheActorIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, actorIdKey)
	return err
}

func (m *defaultActorModel) FindOne(ctx context.Context, id int64) (*Actor, error) {
	actorIdKey := fmt.Sprintf("%s%v", cacheActorIdPrefix, id)
	var resp Actor
	err := m.QueryRowCtx(ctx, &resp, actorIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", actorRows, m.table)
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

func (m *defaultActorModel) Insert(ctx context.Context, data *Actor) (sql.Result, error) {
	actorIdKey := fmt.Sprintf("%s%v", cacheActorIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, actorRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ActorName, data.ActorImg)
	}, actorIdKey)
	return ret, err
}

func (m *defaultActorModel) Update(ctx context.Context, data *Actor) error {
	actorIdKey := fmt.Sprintf("%s%v", cacheActorIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, actorRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ActorName, data.ActorImg, data.Id)
	}, actorIdKey)
	return err
}

func (m *defaultActorModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheActorIdPrefix, primary)
}

func (m *defaultActorModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", actorRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActorModel) tableName() string {
	return m.table
}
