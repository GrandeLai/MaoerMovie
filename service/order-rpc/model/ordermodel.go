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

var _ OrderModel = (*customOrderModel)(nil)

var orderRowsExpectAutoSetButId = strings.Join(stringx.Remove(orderFieldNames, "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`order_time`"), ",")

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllInPageByUserId(ctx context.Context, userId int64, startIndex int64, pageSize int64) ([]*Order, error)
		FindAllPaidInPageByUserId(ctx context.Context, userId int64, startIndex int64, pageSize int64) ([]*Order, error)
		FindAllByShowId(ctx context.Context, showId int64) ([]*Order, error)
		TxInsert(tx *sql.Tx, data *Order) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Order) (sql.Result, error)
		TxDelete(tx *sql.Tx, orderId string) error
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func (m *defaultOrderModel) FindAllInPageByUserId(ctx context.Context, userId int64, startIndex int64, pageSize int64) ([]*Order, error) {
	var resp []*Order
	rowBuilder := m.RowBuilder().Where("user_id = ?", userId)
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

func (m *defaultOrderModel) FindAllPaidInPageByUserId(ctx context.Context, userId int64, startIndex int64, pageSize int64) ([]*Order, error) {
	var resp []*Order
	rowBuilder := m.RowBuilder().Where("user_id = ?", userId).Where("status = ?", int64(1))
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

func (m *defaultOrderModel) FindAllByShowId(ctx context.Context, showId int64) ([]*Order, error) {
	var resp []*Order
	rowBuilder := m.RowBuilder().Where("show_id = ?", showId)
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

func (m *defaultOrderModel) TxInsert(tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSetButId)
	ret, err := tx.Exec(query, data.Uuid, data.CinemaId, data.ShowId, data.FilmId, data.SeatsIds, data.SeatsPosition, data.Price, data.UserId, data.Status)
	return ret, err
}

func (m *defaultOrderModel) TxDelete(tx *sql.Tx, orderId string) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderUuidPrefix, orderId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `uuid` = ?", m.table)
		return tx.Exec(query, orderId)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) TxUpdate(tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("update %s set %s where `uuid` = ?", m.table, orderRowsWithPlaceHolder)
	ret, err := tx.Exec(query, data.CinemaId, data.ShowId, data.FilmId, data.SeatsIds, data.SeatsPosition, data.Price, data.OrderTime, data.UserId, data.Status, data.Uuid)
	return ret, err
}

func (m *defaultOrderModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(orderRows).From(m.table)
}

func (m *defaultOrderModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
