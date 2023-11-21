package pasture

import "gorm.io/gorm"

type PastureSiteDisinfectionRecord struct {
	gorm.Model
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}

type PastureSiteDisinfectionRecordData struct {
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}
