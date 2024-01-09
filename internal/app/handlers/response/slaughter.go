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

type ResSlaughterData struct {
	Data1 slaughter.SlaughterDisinfectHotWaterTempMoniData `json:"data1"`
	Data2 slaughter.SlaughterStunData                      `json:"data2"`
	Data3 slaughter.BleedElectronicData                    `json:"data3"`
	Data4 slaughter.AnalMeatPhMoniData                     `json:"data4"`
	Data5 slaughter.AnalCutWeightData                      `json:"data5"`
	Data6 slaughter.ToNumGermMonData                       `json:"data6"`
	Data7 slaughter.AirNumGermMonData                      `json:"data7"`
}
