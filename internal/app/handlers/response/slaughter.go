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

type ResPreCoolShopData struct {
	ShopInfos []slaughter.PreCoolShopInfo `json:"shop_infos"`
	Count     int64                       `json:"count"`
}

type ResSlaughterShopData struct {
	ShopInfos []slaughter.SlaShopInfo `json:"shop_infos"`
	Count     int64                   `json:"count"`
}

type ResDivisionShopData struct {
	ShopInfos []slaughter.DivShopInfo `json:"shop_infos"`
	Count     int64                   `json:"count"`
}

type ResAcidShopData struct {
	ShopInfos []slaughter.AcidShopInfo `json:"shop_infos"`
	Count     int64                    `json:"count"`
}

type ResFrozenShopData struct {
	ShopInfos []slaughter.FroShopInfo `json:"shop_infos"`
	Count     int64                   `json:"count"`
}

type ResSlaughterWaterQualityData struct {
	Infos []slaughter.SlaughterWaterQualityMonInfo `json:"infos"`
	Count int64                                    `json:"count"`
}

type ResSlaughterStaffUniformData struct {
	Infos []slaughter.StaUniInfo `json:"infos"`
	Count int64
}

type ResSlaughterLightRecord struct {
	Infos []slaughter.SlaughterLightRecordInfo `json:"infos"`
	Count int64
}
