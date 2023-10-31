package coldchain

import "gorm.io/gorm"

type ColdchainProcedure struct {
	gorm.Model
	CcPID string `gorm:"not null; unique; type:varchar(256)" json:"cc_pid"`
}
