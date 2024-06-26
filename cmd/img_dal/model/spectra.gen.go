// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSpectra = "spectra"

// Spectra mapped from table <spectra>
type Spectra struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Data       string    `gorm:"column:data;not null" json:"data"`
	UploadTime time.Time `gorm:"column:upload_time;default:CURRENT_TIMESTAMP" json:"upload_time"`
}

// TableName Spectra's table name
func (*Spectra) TableName() string {
	return TableNameSpectra
}
