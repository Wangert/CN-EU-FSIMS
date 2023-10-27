package models

import "gorm.io/gorm"

// 用户表
type FSIMSUser struct {
	gorm.Model
	UUID         string `gorm:"not null" json:"uuid"`
	Name         string `gorm:"not null" json:"name"`
	Account      string `gorm:"not null" json:"account"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	Type         string `json:"type"`
	Status       string `gorm:"not null" json:"status"`
	Company      string `gorm:"not null" json:"company"`
	Phone        string `gorm:"not null" json:"phone"`
}

type User struct {
	ID        int
	Username  string
	Password  string
	Company   string
	Usertype  int
	Telephone string
	WarnNum   int
	Login     bool
	Place     string
}
