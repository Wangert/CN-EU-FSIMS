package pasture

import "gorm.io/gorm"

type PastureFodderPhysicalHazard struct {
	gorm.Model
	Mercury                  float64 `json:"mercury"`                   // 汞
	Cadmium                  float64 `json:"cadmium"`                   // 镉
	Lead                     float64 `json:"lead"`                      // 铅
	Chromium                 float64 `json:"chromium"`                  // 铬
	Arsenic                  float64 `json:"arsenic"`                   // 砷
	Fluorine                 float64 `json:"fluorine"`                  // 氟
	Nitrite                  float64 `json:"nitrite"`                   // 亚硝酸盐（以NaNO2计）
	PolychlorinatedBiphenyls float64 `json:"polychlorinated_biphenyls"` // 多氯联苯
	Sixsixsix                float64 `json:"sixsixsix"`                 // 六六六
	DDT                      float64 `json:"ddt"`                       // 滴滴涕
	Bexachlorobenzene        float64 `json:"bexachlorobenzene"`         // 六氯苯
	PastureFodderID          uint    `json:"pasture_fodder_id"`
}

type PastureFodderPhysicalHazardData struct {
	Mercury                  float64 `json:"mercury"`                   // 汞
	Cadmium                  float64 `json:"cadmium"`                   // 镉
	Lead                     float64 `json:"lead"`                      // 铅
	Chromium                 float64 `json:"chromium"`                  // 铬
	Arsenic                  float64 `json:"arsenic"`                   // 砷
	Fluorine                 float64 `json:"fluorine"`                  // 氟
	Nitrite                  float64 `json:"nitrite"`                   // 亚硝酸盐（以NaNO2计）
	PolychlorinatedBiphenyls float64 `json:"polychlorinated_biphenyls"` // 多氯联苯
	Sixsixsix                float64 `json:"sixsixsix"`                 // 六六六
	DDT                      float64 `json:"ddt"`                       // 滴滴涕
	Bexachlorobenzene        float64 `json:"bexachlorobenzene"`         // 六氯苯
	PastureFodderID          uint    `json:"pasture_fodder_id"`
}
