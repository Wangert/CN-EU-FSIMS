package pasture

import "gorm.io/gorm"

type TruckDisinfectionRecord struct {
	gorm.Model
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}
