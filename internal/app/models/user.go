package models

import (
	"CN-EU-FSIMS/internal/app/handlers/response"
	"gorm.io/gorm"
)

// user table
type FSIMSUser struct {
	gorm.Model
	UUID         string `gorm:"not null" json:"uuid"`
	Name         string `gorm:"not null" json:"name"`
	Account      string `gorm:"not null" json:"account"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	Type         int    `json:"type"`
	Role         string `json:"role"`
	Status       int    `gorm:"not null" json:"status"`
	Company      string `gorm:"not null" json:"company"`
	Phone        string `gorm:"not null" json:"phone"`
}

func FsimsUserToResUser(fsimsUser *FSIMSUser) response.ResUser {
	return response.ResUser{
		UUID:    fsimsUser.UUID,
		Name:    fsimsUser.Name,
		Account: fsimsUser.Account,
		Type:    fsimsUser.Type,
		Role:    fsimsUser.Role,
		Status:  fsimsUser.Status,
		Company: fsimsUser.Company,
		Phone:   fsimsUser.Phone,
	}
}
