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
	State         int        `gorm:"not null" json:"state"`
	Operator      string     `gorm:"type:varchar(100)" json:"operator"`
	ReceiveTime   *time.Time `gorm:"not null" json:"receive_time"`
	ConfirmTime   *time.Time `json:"confirm_time"`
	PacNumber     string     `gorm:"not null; type:varchar(256)" json:"pac_number"`
}

type PackageReceiveRecordInfo struct {
	ProductNumber string `json:"product_number"`
	PID           string `json:"pid"`
	SourceNumber  string `json:"source_number"`
	SourceName    string `json:"source_name"`
	State         int    `json:"state"`
	Operator      string `json:"operator"`
	ReceiveTime   string `json:"receive_time"`
	ConfirmTime   string `json:"confirm_time"`
	PacNumber     string `json:"pac_number"`
}

func ToPackageReceiveRecordInfo(record *PackageReceiveRecord) PackageReceiveRecordInfo {
	receivetime := ""
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// receivetime := record.ReceiveTime.In(loc)
	// confirmtime := record.ConfirmTime.In(loc)
	if record.ReceiveTime != nil {
		stime := record.ReceiveTime.In(loc)
		receivetime = stime.Format("2006-01-02 15:04:05")
	}
	confirmtime := ""
	if record.ConfirmTime != nil {
		etime := record.ConfirmTime.In(loc)
		confirmtime = etime.Format("2006-01-02 15:04:05")
	}
	return PackageReceiveRecordInfo{
		ProductNumber: record.ProductNumber,
		PID:           record.PID,
		SourceNumber:  record.SourceNumber,
		SourceName:    record.SourceName,
		State:         record.State,
		Operator:      record.Operator,
		ReceiveTime:   receivetime,
		ConfirmTime:   confirmtime,
		PacNumber:     record.PacNumber,
	}
}
