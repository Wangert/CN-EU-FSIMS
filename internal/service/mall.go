package service

import (
	"CN-EU-FSIMS/internal/app/models/product"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
)

func GetMallGoods(mallNumber string) ([]product.MallGoodInfo, int64, error) {
	q := query.MallGood
	goods, err := q.WithContext(context.Background()).Where(q.MallNumber.Eq(mallNumber)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(goods)
	records := make([]product.MallGoodInfo, count)
	for i, g := range goods {
		records[i] = product.ToMallGoodInfo(g)
	}

	return records, int64(count), nil
}

func GetAllMallGoods() ([]product.MallGoodInfo, int64, error) {
	q := query.MallGood
	goods, err := q.WithContext(context.Background()).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(goods)
	records := make([]product.MallGoodInfo, count)
	for i, g := range goods {
		records[i] = product.ToMallGoodInfo(g)
	}

	return records, int64(count), nil
}
