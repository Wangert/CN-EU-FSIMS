package response

import (
	"CN-EU-FSIMS/internal/app/models/slaughter"
	"CN-EU-FSIMS/internal/app/models/warehouse"
)

type ResSlaughterReceiveRecords struct {
	Records []warehouse.SlaughterReceiveRecordInfo `json:"records"`
	Count   int64                                  `json:"count"`
}

type ResSlaughterBatches struct {
	Records []slaughter.SlaughterBatchInfo `json:"records"`
	Count   int64                          `json:"count"`
}

type ResSlaughterWarehouseRecords struct {
	Records []warehouse.SlaughterWarehouseInfo `json:"records"`
	Count   int64                              `json:"count"`
}

type ResEndSlaughter struct {
	Checkcode   string   `json:"checkcode"`
	ProductsNum []string `json:"products_num"`
	Count       int64    `json:"count"`
}
