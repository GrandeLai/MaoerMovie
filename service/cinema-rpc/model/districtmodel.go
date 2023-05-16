package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DistrictModel = (*customDistrictModel)(nil)

type (
	// DistrictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDistrictModel.
	DistrictModel interface {
		districtModel
		FindAllByCityName(ctx context.Context, cityName string) ([]*District, error)
	}

	customDistrictModel struct {
		*defaultDistrictModel
	}
)

// NewDistrictModel returns a model for the database table.
func NewDistrictModel(conn sqlx.SqlConn, c cache.CacheConf) DistrictModel {
	return &customDistrictModel{
		defaultDistrictModel: newDistrictModel(conn, c),
	}
}

func (m *defaultDistrictModel) FindAllByCityName(ctx context.Context, cityName string) ([]*District, error) {
	//var resp []*District
	//rowBuilder := m.RowBuilder().Where("city_name = ?", cityName)
	//query, values, err := rowBuilder.ToSql()
	//err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	//switch err {
	//case nil:
	//	return resp, nil
	//case sqlx.ErrNotFound:
	//	return nil, ErrNotFound
	//default:
	//	return nil, err
	//}
	var resp []*District
	rowBuilder := m.RowBuilder().Where("city_name = ?", cityName)
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

func (m *defaultDistrictModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(districtRows).From(m.table)
}
