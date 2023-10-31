package pasture

import "gorm.io/gorm"

type PastureFodderPhysicalHazard struct {
	gorm.Model
	Mercury                  float32 `json:"mercury"`                   // 汞
	Cadmium                  float32 `json:"cadmium"`                   // 镉
	Lead                     float32 `json:"lead"`                      // 铅
	Chromium                 float32 `json:"chromium"`                  // 铬
	Arsenic                  float32 `json:"arsenic"`                   // 砷
	Fluorine                 float32 `json:"fluorine"`                  // 氟
	Nitrite                  float32 `json:"nitrite"`                   // 亚硝酸盐（以NaNO2计）
	PolychlorinatedBiphenyls float32 `json:"polychlorinated_biphenyls"` // 多氯联苯
	Sixsixsix                float32 `json:"sixsixsix"`                 // 六六六
	DDT                      float32 `json:"ddt"`                       // 滴滴涕
	Bexachlorobenzene        float32 `json:"bexachlorobenzene"`         // 六氯苯
	PastureFodderID          uint    `json:"pasture_fodder_id"`
}
