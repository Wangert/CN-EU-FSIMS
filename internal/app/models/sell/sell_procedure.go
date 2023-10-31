package sell

import "gorm.io/gorm"

type SellProcedure struct {
	gorm.Model
	SePID string `gorm:"not null; unique; type:varchar(256)" json:"se_pid"`
}
