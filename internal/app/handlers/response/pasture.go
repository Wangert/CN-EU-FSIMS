package response

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
	"CN-EU-FSIMS/internal/app/models/warehouse"
)

type ResFeedingRecords struct {
	FeedingBatches []pasture.FeedingBatchInfo `json:"feeding_batches"`
	Count          int64                      `json:"count"`
}

type ResEndFeeding struct {
	Checkcode string   `json:"checkcode"`
	CowsNum   []string `json:"cows_num"`
	Count     int64    `json:"count"`
}

type ResWarehouseInfos struct {
	PastureWarehouses []warehouse.PastureWarehouseInfo `json:"pasture_warehouses"`
	Count             int64                            `json:"count"`
}
