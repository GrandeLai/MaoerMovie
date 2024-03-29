package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShowModel = (*customShowModel)(nil)

type (
	// ShowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShowModel.
	ShowModel interface {
		showModel
		FindByFactors(ctx context.Context, cinemaId int64, filmId int64, date string) ([]*Show, error)
		TxUpdateSurplusNumWithLock(tx *sql.Tx, num int64, showId int64) (sql.Result, error)
	}

	customShowModel struct {
		*defaultShowModel
	}
)

// NewShowModel returns a model for the database table.
func NewShowModel(conn sqlx.SqlConn, c cache.CacheConf) ShowModel {
	return &customShowModel{
		defaultShowModel: newShowModel(conn, c),
	}
}
func (m *defaultShowModel) FindByFactors(ctx context.Context, cinemaId int64, filmId int64, date string) ([]*Show, error) {
	var resp []*Show
	rowBuilder := m.RowBuilder()
	query, values, err := rowBuilder.Where("cinema_id=?", cinemaId).Where("film_id=?", filmId).Where("DATE(date) = ?", date).ToSql()
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

func (m *defaultShowModel) TxUpdateSurplusNumWithLock(tx *sql.Tx, num int64, showId int64) (sql.Result, error) {
	showIdKey := fmt.Sprintf("%s%v", cacheShowIdPrefix, showId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set surplus_num = surplus_num + ?  where `id` = ? and `surplus_num` > 0", m.table)
		return tx.Exec(query, num, showId)
	}, showIdKey)
	return ret, err
}

func (m *defaultShowModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(showRows).From(m.table)
}
