package pasture

import "gorm.io/gorm"

type FloorBeddingBiohazard struct {
	gorm.Model
	ColiformBacteria uint `json:"coliform_bacteria"` // 总大肠菌群
	TotalBacteria    uint `json:"total_bacteria"`    // 总细菌数
	TotalMold        uint `json:"total_mold"`        // 总霉菌数
	AflatoxinB1      uint `json:"aflatoxin_b1"`      // 黄曲霉素B1
	Salmonella       uint `json:"salmonella"`        // 沙门菌数
	FloorBeddingID   uint `json:"floor_bedding_id"`
}
