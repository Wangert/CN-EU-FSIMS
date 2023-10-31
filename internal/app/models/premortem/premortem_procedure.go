package premortem

import "gorm.io/gorm"

type PremortemProcedure struct {
	gorm.Model
	PremPID string `gorm:"not null; unique; type:varchar(256)" json:"prem_pid"`
}
