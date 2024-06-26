// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLymphNode = "lymph_nodes"

// LymphNode mapped from table <lymph_nodes>
type LymphNode struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`               // 自增主键
	Grade      string    `gorm:"column:grade;not null;comment:淋巴结等级" json:"grade"`                             // 淋巴结等级
	ImgPath    string    `gorm:"column:img_path;not null;comment:图片路径" json:"img_path"`                        // 图片路径
	UploadTime time.Time `gorm:"column:upload_time;default:CURRENT_TIMESTAMP;comment:上传时间" json:"upload_time"` // 上传时间
}

// TableName LymphNode's table name
func (*LymphNode) TableName() string {
	return TableNameLymphNode
}
