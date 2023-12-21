package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type PackReceive struct {
	gorm.Model
	ProductNumber    string    `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	ProductPID       string    `gorm:"not null" json:"product_pid"`
	ProductType      string    `gorm:"not null; type:varchar(100)" json:"product_type"`
	Operator         string    `gorm:"not null; type:varchar(100)" json:"operator"`
	ReceiveTimestamp time.Time `gorm:"not null" json:"receive_timestamp"`
}

type PackageReceiveRecord struct {
	gorm.Model
	ProductNumber string     `gorm:"not null; unique; type:varchar(256)" json:"product_number"`
	PID           string     `gorm:"not null" json:"pid"`
	SourceNumber  string     `gorm:"not null" json:"source_number"`
	SourceName    string     `gorm:"not null" json:"source_name"`
	Operator      string     `gorm:"type:varchar(100)" json:"operator"`
	ReceiveTime   time.Time  `gorm:"not null" json:"receive_time"`
	ConfirmTime   *time.Time `json:"confirm_time"`
	PacNumber     string     `gorm:"not null; type:varchar(256)" json:"pac_number"`
}

type PackageReceiveRecordInfo struct {
	ProductNumber string     `json:"product_number"`
	PID           string     `json:"pid"`
	SourceNumber  string     `json:"source_number"`
	SourceName    string     `json:"source_name"`
	Operator      string     `json:"operator"`
	ReceiveTime   time.Time  `json:"receive_time"`
	ConfirmTime   *time.Time `json:"confirm_time"`
	PacNumber     string     `json:"pac_number"`
}

func ToPackageReceiveRecordInfo(record PackageReceiveRecord) PackageReceiveRecordInfo {
	return PackageReceiveRecordInfo{
		ProductNumber: record.ProductNumber,
		PID:           record.PID,
		SourceNumber:  record.SourceNumber,
		SourceName:    record.SourceName,
		Operator:      record.Operator,
		ReceiveTime:   record.ReceiveTime,
		ConfirmTime:   record.ConfirmTime,
		PacNumber:     record.PacNumber,
	}
}
