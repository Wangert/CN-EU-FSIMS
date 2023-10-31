package pasture

import "gorm.io/gorm"

type WasteDischarge struct {
	gorm.Model
	WaterFivedayBOD          float32 `json:"water_fiveday_bod"`            // 水 五日生化需氧量
	WaterChemicalOxygen      float32 `json:"water_chemical_oxygen"`        // 水 化学需氧量
	WaterAmmoniaNitrogen     float32 `json:"water_ammonia_nitrogen"`       // 水 氨氮
	WaterPhosphorus          float32 `json:"water_phosphorus"`             // 水 磷
	WaterSuspendedMatter     float32 `json:"water_suspended_matter"`       // 水 悬浮物
	WaterFecalColiform       uint    `json:"fecal_coliform"`               // 水 粪大肠菌群
	WaterAO                  float32 `json:"water_ao"`                     // 水 蛔虫卵
	WasteSlagFecalColiform   uint    `json:"waste_slag_fecal_coliform"`    // 废渣 粪大肠菌群
	WasteSlagAOMortalityRate float32 `json:"waste_slag_ao_mortality_rate"` // 废渣 蛔虫卵死亡率
	O3Concentration          float32 `json:"o3_concentration"`             // 臭氧浓度
	PasPID                   string  `json:"pas_pid"`
}
