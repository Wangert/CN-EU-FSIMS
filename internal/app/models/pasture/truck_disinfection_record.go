package pasture

import "gorm.io/gorm"

type PastureTruckDisinfectionRecord struct {
	gorm.Model
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}

type PastureTruckDisinfectionRecordData struct {
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}
