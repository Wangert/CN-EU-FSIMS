package slaughter

import (
	"gorm.io/gorm"
	"time"
)

type ChiledFreshIndex struct {
	//冷鲜肉相关指标
	gorm.Model
	TimeRecordAt        time.Time           `json:"time_record_at"`                                                            //记录时间
	SlaughterPaGerm     SlaughterPaGerm     `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"slaughter_pa_germ"`     ////致病菌监测
	SlaughterOtherIndex SlaughterOtherIndex `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"slaughter_other_index"` //其他指标
	SlaughterDrugsResi  SlaughterDrugsResi  `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"slaughter_drugs_resi"`  //农兽药残留
	SlaughterGps        SlaughterGps        `gorm:"foreignKey:ChiledFreshIndexID; references:ID" json:"slaughter_gps"`
	TranTemp            float32             `json:"tran_temp"` //运输温度
}

type SlaughterGps struct {
	//GPS
	gorm.Model
	ChiledFreshIndexID *uint  `json:"chiled_fresh_index_id"` //外键
	Gps1               string `json:"gps_1"`                 //gps坐标
	Gps2               string `json:"gps_2"`                 //时间
}

type SlaughterPaGerm struct {
	//致病菌监测
	gorm.Model
	ChiledFreshIndexID *uint   `json:"chiled_fresh_index_id"`
	PaGerm1            float32 `json:"pa_germ_1"` //光谱监测沙门
	PaGerm2            float32 `json:"pa_germ_2"` //STEC
	PaGerm3            float32 `json:"pa_germ_3"` //李斯特
	PaGerm4            float32 `json:"pa_germ_4"` //阪崎杆菌
	PaGerm5            float32 `json:"pa_germ_5"` //金葡
}

type SlaughterOtherIndex struct {
	//其他指标
	gorm.Model
	ChiledFreshIndexID *uint   `json:"chiled_fresh_index_id"`
	OtherIndex1        float32 `json:"other_index_1"` //菌落总数（光谱）
	OtherIndex2        float32 `json:"other_index_2"` //挥发性盐基氮（光谱）
	OtherIndex3        float32 `json:"other_index_3"` //大肠菌群（光谱）
	OtherIndex4        float32 `json:"other_index_4"` //假单胞
	OtherIndex5        float32 `json:"other_index_5"` //乳酸菌（光谱）
	OtherIndex6        float32 `json:"other_index_6"` //铅
	OtherIndex7        float32 `json:"other_index_7"` //无机砷
	OtherIndex8        float32 `json:"other_index_8"` //牛肉水分
	OtherIndex9        float32 `json:"other_index_9"` //牛肉蛋白质
}

type SlaughterDrugsResi struct {
	//农兽药残留
	gorm.Model
	ChiledFreshIndexID *uint   `json:"chiled_fresh_index_id"`
	DrugsResi1         float32 `json:"drugs_resi_1"` //恩诺沙星
	DrugsResi2         float32 `json:"drugs_resi_2"` //庆大霉素
	DrugsResi3         float32 `json:"drugs_resi_3"` //青霉素
	DrugsResi4         float32 `json:"drugs_resi_4"` //氯氢吡啶
}

type ChiledFreshIndexInfo struct {
	//冷鲜肉相关指标
	TimeRecordAt        string                  `json:"time_record_at"`        //记录时间
	SlaughterPaGerm     SlaughterPaGermInfo     `json:"slaughter_pa_germ"`     ////致病菌监测
	SlaughterOtherIndex SlaughterOtherIndexInfo `json:"slaughter_other_index"` //其他指标
	SlaughterDrugsResi  SlaughterDrugsResiInfo  `json:"slaughter_drugs_resi"`  //农兽药残留
	SlaughterGps        SlaughterGpsInfo        `json:"slaughter_gps"`
	TranTemp            float32                 `json:"tran_temp"` //运输温度
}

type SlaughterGpsInfo struct {
	//GPS
	Gps1 string `json:"gps_1"` //gps坐标
	Gps2 string `json:"gps_2"` //时间
}

type SlaughterPaGermInfo struct {
	//致病菌监测
	PaGerm1 float32 `json:"pa_germ_1"` //光谱监测沙门
	PaGerm2 float32 `json:"pa_germ_2"` //STEC
	PaGerm3 float32 `json:"pa_germ_3"` //李斯特
	PaGerm4 float32 `json:"pa_germ_4"` //阪崎杆菌
	PaGerm5 float32 `json:"pa_germ_5"` //金葡
}

type SlaughterOtherIndexInfo struct {
	//其他指标
	OtherIndex1 float32 `json:"other_index_1"` //菌落总数（光谱）
	OtherIndex2 float32 `json:"other_index_2"` //挥发性盐基氮（光谱）
	OtherIndex3 float32 `json:"other_index_3"` //大肠菌群（光谱）
	OtherIndex4 float32 `json:"other_index_4"` //假单胞
	OtherIndex5 float32 `json:"other_index_5"` //乳酸菌（光谱）
	OtherIndex6 float32 `json:"other_index_6"` //铅
	OtherIndex7 float32 `json:"other_index_7"` //无机砷
	OtherIndex8 float32 `json:"other_index_8"` //牛肉水分
	OtherIndex9 float32 `json:"other_index_9"` //牛肉蛋白质
}

type SlaughterDrugsResiInfo struct {
	//农兽药残留
	DrugsResi1 float32 `json:"drugs_resi_1"` //恩诺沙星
	DrugsResi2 float32 `json:"drugs_resi_2"` //庆大霉素
	DrugsResi3 float32 `json:"drugs_resi_3"` //青霉素
	DrugsResi4 float32 `json:"drugs_resi_4"` //氯氢吡啶
}

func ToChiledFreshIndexInfo(index *ChiledFreshIndex) ChiledFreshIndexInfo {

	germ := SlaughterPaGermInfo{
		PaGerm1: index.SlaughterPaGerm.PaGerm1,
		PaGerm2: index.SlaughterPaGerm.PaGerm2,
		PaGerm3: index.SlaughterPaGerm.PaGerm3,
		PaGerm4: index.SlaughterPaGerm.PaGerm4,
		PaGerm5: index.SlaughterPaGerm.PaGerm5,
	}

	otherIndex := SlaughterOtherIndexInfo{
		OtherIndex1: index.SlaughterOtherIndex.OtherIndex1,
		OtherIndex2: index.SlaughterOtherIndex.OtherIndex2,
		OtherIndex3: index.SlaughterOtherIndex.OtherIndex3,
		OtherIndex4: index.SlaughterOtherIndex.OtherIndex4,
		OtherIndex5: index.SlaughterOtherIndex.OtherIndex5,
		OtherIndex6: index.SlaughterOtherIndex.OtherIndex6,
		OtherIndex7: index.SlaughterOtherIndex.OtherIndex7,
		OtherIndex8: index.SlaughterOtherIndex.OtherIndex8,
		OtherIndex9: index.SlaughterOtherIndex.OtherIndex9,
	}

	resi := SlaughterDrugsResiInfo{
		DrugsResi1: index.SlaughterDrugsResi.DrugsResi1,
		DrugsResi2: index.SlaughterDrugsResi.DrugsResi2,
		DrugsResi3: index.SlaughterDrugsResi.DrugsResi3,
		DrugsResi4: index.SlaughterDrugsResi.DrugsResi4,
	}

	gps := SlaughterGpsInfo{
		Gps1: index.SlaughterGps.Gps1,
		Gps2: index.SlaughterGps.Gps2,
	}

	return ChiledFreshIndexInfo{
		TimeRecordAt:        index.TimeRecordAt.Format("2006-01-02 15:04:05"),
		SlaughterPaGerm:     germ,
		SlaughterOtherIndex: otherIndex,
		SlaughterDrugsResi:  resi,
		SlaughterGps:        gps,
		TranTemp:            index.TranTemp,
	}
}
