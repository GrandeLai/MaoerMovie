package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CinemaModel = (*customCinemaModel)(nil)

type (
	// CinemaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCinemaModel.
	CinemaModel interface {
		cinemaModel
		FindByFactors(ctx context.Context, brandId int64, hallTypeId string, districtId int64, startIndex int64, pageSize int64) ([]*Cinema, error)
	}

	customCinemaModel struct {
		*defaultCinemaModel
	}
)

// NewCinemaModel returns a model for the database table.
func NewCinemaModel(conn sqlx.SqlConn, c cache.CacheConf) CinemaModel {
	return &customCinemaModel{
		defaultCinemaModel: newCinemaModel(conn, c),
	}
}

func (m *defaultCinemaModel) FindByFactors(ctx context.Context, brandId int64, hallTypeId string, districtId int64, startIndex int64, pageSize int64) ([]*Cinema, error) {
	var resp []*Cinema
	rowBuilder := m.RowBuilder()
	if brandId != 0 {
		rowBuilder = rowBuilder.Where("brand_id = ?", brandId)
	}
	if hallTypeId != "" {
		rowBuilder = rowBuilder.Where("hall_ids LIKE ?", "%/"+hallTypeId+"/%")
	}
	if districtId != 0 {
		rowBuilder = rowBuilder.Where("district_id = ?", districtId)
	}
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

func (m *defaultCinemaModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(cinemaRows).From(m.table)
}

func (m *defaultCinemaModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("count(" + field + ")").From(m.table)
}
