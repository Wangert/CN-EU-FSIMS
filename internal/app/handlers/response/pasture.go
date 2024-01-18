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

type ResFeedHeavyMetalRecords struct {
	Count                 int64                               `json:"count"`
	FeedHeavyMetalRecords []pasture.PastureFeedHeavyMetalInfo `json:"feed_heavy_metal_records"`
}

type ResFeedMycotoxinsRecords struct {
	Count                 int64                               `json:"count"`
	FeedMycotoxinsRecords []pasture.PastureFeedMycotoxinsInfo `json:"feed_mycotoxins_records"`
}

type ResPastureWaterRecords struct {
	Count               int64                            `json:"count"`
	PastureWaterRecords []pasture.PastureWaterRecordInfo `json:"pasture_water_records"`
}

type ResPastureDisinfectionRecords struct {
	Count                      int64                                   `json:"count"`
	PastureDisinfectionRecords []pasture.PastureDisinfectionRecordInfo `json:"pasture_disinfection_records"`
}

type ResPastureWastedWaterIndexRecords struct {
	Count                          int64                                 `json:"count"`
	PastureWastedWaterIndexRecords []pasture.PastureWastedWaterIndexInfo `json:"pasture_wasted_water_index_records"`
}

type ResPasturePaddingRequireRecords struct {
	Count                        int64                               `json:"count"`
	PasturePaddingRequireRecords []pasture.PasturePaddingRequireInfo `json:"pasture_padding_require_records"`
}

type ResPastureBasicEnvironmentRecords struct {
	Count                          int64                                 `json:"count"`
	PastureBasicEnvironmentRecords []pasture.PastureBasicEnvironmentInfo `json:"pasture_basic_environment_records"`
}

type ResPastureCowHouseRecords struct {
	Count                  int64                  `json:"count"`
	PastureCowHouseRecords []pasture.CowHouseInfo `json:"pasture_cow_house_records"`
}

type ResPastureAreaRecords struct {
	Count              int64                     `json:"count"`
	PastureAreaRecords []pasture.PastureAreaInfo `json:"pasture_area_records"`
}

type ResPastureBufferRecords struct {
	Count                int64                       `json:"count"`
	PastureBufferRecords []pasture.PastureBufferInfo `json:"pasture_buffer_records"`
}

type ResWasteResidueOdor struct {
	ResWasteResidueOdor1 float64 `json:"res_waste_residue_odor_1"` //废渣正常总量
	ResWasteResidueOdor2 float64 `json:"res_waste_residue_odor_2"` //正常较昨日新增
	ResWasteResidueOdor3 float64 `json:"res_waste_residue_odor_3"` //废渣超标总量
	ResWasteResidueOdor4 float64 `json:"res_waste_residue_odor_4"` //废渣超标较昨日新增
	ResWasteResidueOdor5 float64 `json:"res_waste_residue_odor_5"` //恶臭污染物正常总量
	ResWasteResidueOdor6 float64 `json:"res_waste_residue_odor_6"` //恶臭污染物正常较昨日新增
	ResWasteResidueOdor7 float64 `json:"res_waste_residue_odor_7"` //odr超标总量
	ResWasteResidueOdor8 float64 `json:"res_waste_residue_odor_8"` //odr超标较昨日新增
}
