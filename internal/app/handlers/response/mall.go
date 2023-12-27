package response

import (
	"CN-EU-FSIMS/internal/app/models/product"
)

type ResMallGoods struct {
	Records []product.MallGoodInfo `json:"records"`
	Count   int64                  `json:"count"`
}
