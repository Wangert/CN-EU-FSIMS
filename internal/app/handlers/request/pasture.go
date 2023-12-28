package request

import (
	"CN-EU-FSIMS/internal/app/models/pasture"
)

type ReqAddCow struct {
	Age         int     `json:"age" form:"age"`
	Weight      float64 `json:"weight" form:"weight"`
	HouseNumber string  `json:"house_number" form:"house_number"`
}

type ReqNewFeedingBatch struct {
	HouseNumber string   `json:"house_number" form:"house_number"`
	Worker      string   `json:"worker" form:"worker"`
	PrePID      string   `json:"pre_pid" form:"pre_pid"`
	CowNumbers  []string `json:"cow_numbers" form:"cow_numbers"`
}

type ReqEndFeeding struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	PM10        float32 `json:"pm_10" form:"pm_10"`   // PM10
	TSP         float32 `json:"tsp" form:"tsp"`       // TSP
	Stench      uint    `json:"stench" form:"stench"` // 恶臭稀释程度
}

type ReqSendToSlaughter struct {
	CowNumber            string `json:"cow_number" form:"cow_number"`
	Operator             string `json:"operator" form:"operator"`
	SlaughterHouseNumber string `json:"slaughter_house_number" form:"slaughter_house_number"`
}

type ReqAddPastureFeedHeavyMetal struct {
	HouseNumber string                    `json:"house_number"`
	Ass         pasture.PastureFeedAsInfo `json:"aas"` //砷元素
	Pb          pasture.PastureFeedPbInfo `json:"pb"`  //铅元素
	Cd          pasture.PastureFeedCdInfo `json:"cd"`  //镉元素
	Cr          pasture.PastureFeedCrInfo `json:"cr"`  //铬元素
}

type ReqAddPastureFeedCass struct {
	////饲料中真菌毒素、农兽药残留
	HouseNumber string               `json:"house_number"`
	Afb1        pasture.Afb1Info     `json:"afb_1"`       //黄曲霉毒素B1
	Don         pasture.DonInfo      `json:"don"`         //玉米赤霉烯酮
	T2toxin     pasture.T2toxinInfo  `json:"t2Toxin"`     //脱氧雪腐镰刀菌烯醇（呕吐毒素）
	T2VomZea    pasture.T2VomZeaInfo `json:"t_2_vom_zea"` //T-2毒素 伏马毒素 赭曲霉毒素A
}

type ReqAddPastureWaterRecord struct {
	HouseNumber string                        `json:"house_number"` //时间记录
	OapGci      pasture.PastureOapGciInfo     `json:"oap_gci"`      //感官性状和一般化学指标
	ToxIndex    pasture.PastureToxIndexInfo   `json:"tox_index"`    //毒理性指标
	MicroIndex  pasture.PastureMicroIndexInfo `json:"micro_index"`  //微生物指标
}
