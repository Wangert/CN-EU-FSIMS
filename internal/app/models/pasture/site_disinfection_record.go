package pasture

import "gorm.io/gorm"

type SiteDisinfectionRecord struct {
	gorm.Model
	Method        string `json:"method"`
	Concentration uint   `json:"concentration"`
	Duration      uint64 `json:"duration"`
}
