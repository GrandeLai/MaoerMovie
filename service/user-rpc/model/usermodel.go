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

var _ UserModel = (*customUserModel)(nil)

var userRowsExpectAutoSetButId = strings.Join(stringx.Remove(userFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		CountByPhone(ctx context.Context, phone string) (int, error)
		CountByEmail(ctx context.Context, email string) (int, error)
		InsertWithNewId(ctx context.Context, data *User) (sql.Result, error)
		FindAll(ctx context.Context) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) CountByPhone(ctx context.Context, phone string) (int, error) {
	rowBuilder := m.CountBuilder("id")
	query, values, err := rowBuilder.Where("phone = ?", phone).ToSql()
	if err != nil {
		return 0, err
	}
	var resp int
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserModel) CountByEmail(ctx context.Context, email string) (int, error) {
	rowBuilder := m.CountBuilder("id")
	query, values, err := rowBuilder.Where("email = ?", email).ToSql()
	if err != nil {
		return 0, err
	}
	var resp int
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserModel) InsertWithNewId(ctx context.Context, data *User) (sql.Result, error) {
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userMobileKey := fmt.Sprintf("%s%v", cacheUserPhonePrefix, data.Phone)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?,?)", m.table, userRowsExpectAutoSetButId)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.Gender, data.Phone, data.Password, data.Email, data.Status, data.AvatarUrl)
	}, userEmailKey, userIdKey, userMobileKey)
	return ret, err
}

func (m *defaultUserModel) FindAll(ctx context.Context) ([]*User, error) {
	var resp []*User
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("status = ?", 0).ToSql()
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

func (m *defaultUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRows).From(m.table)
}
func (m *defaultUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
