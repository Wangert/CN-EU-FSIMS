package pasture

import "gorm.io/gorm"

type PastureFodderBiohazard struct {
	gorm.Model
	AflatoxinB1       uint `json:"aflatoxin_b1"`      // 黄曲霉素B1
	OchratoxinA       uint `json:"ochratoxin_a"`      // 赭曲霉毒素A
	Zearalenone       uint `json:"zearalenone"`       // 玉米赤霉烯酮
	Vomitoxin         uint `json:"vomitoxin"`         // 呕吐毒素
	T2                uint `json:"t_2"`               // T-2毒素
	B1B2              uint `json:"b1_b2"`             // 伏马毒素（B1＋B2）
	Cyanide           uint `json:"cyanide"`           // 氰化物（以HCN计）
	FreeGossypol      uint `json:"free_gossypol"`     // 游离棉酚
	Isothiocyanate    uint `json:"isothiocyanate"`    // 异硫氰酸酯
	Oxazolidinethione uint `json:"oxazolidinethione"` // 噁唑烷硫酮
	TotalBacteria     uint `json:"total_bacteria"`    // 总细菌数
	TotalMold         uint `json:"total_mold"`        // 总霉菌数
	Salmonella        uint `json:"salmonella"`        // 沙门菌数
	PastureFodderID   uint `json:"pasture_fodder_id"`
}
