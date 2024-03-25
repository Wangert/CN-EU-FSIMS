package models

import (
	"encoding/json"
	"time"
)

// ObjectDetection 结构体映射了JSON对象的结构
type ObjectDetection struct {
	Name       string  `json:"name"`
	Xmax       float64 `json:"xmax"`
	Xmin       float64 `json:"xmin"`
	Ymax       float64 `json:"ymax"`
	Ymin       float64 `json:"ymin"`
	Class      int     `json:"class"`
	Confidence float64 `json:"confidence"`
}

type Spectra struct {
	Id         uint            `gorm:"primarykey" json:"id"`
	Data       json.RawMessage `gorm:"not null" json:"data"`
	UploadTime time.Time       `gorm:"not null" json:"upload_time"`
}

type UploadImgs struct {
	Id         uint            `gorm:"primarykey" json:"id"`
	Filename   string          `gorm:"not null" json:"filename"`
	Source     string          `gorm:"not null" json:"source"`
	UploadTime time.Time       `gorm:"not null" json:"upload_time"`
	Result     json.RawMessage `gorm:"not null" json:"result"`
}

type UploadImgsInfo struct {
	Filename   string          `json:"filename"`
	Source     string          `json:"source"`
	UploadTime string          `json:"upload_time"`
	Result     json.RawMessage `json:"result"`
}

func ToUploadImgsInfo(ui *UploadImgs) UploadImgsInfo {
	return UploadImgsInfo{
		Filename:   ui.Filename,
		Source:     ui.Source,
		UploadTime: ui.UploadTime.Format("2006-01-02 15:04:05"),
		Result:     ui.Result,
	}
}
