package _package

import "gorm.io/gorm"

type PackageProcedure struct {
	gorm.Model
	PacPID string `gorm:"not null; unique; type:varchar(256)" json:"pac_pid"`
}
