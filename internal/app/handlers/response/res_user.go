package response

import "CN-EU-FSIMS/internal/app/models"

type ResUsers struct {
	Users []models.UserInfo `json:"users"`
	Count int64             `json:"count"`
}

type ResLogin struct {
	UUID        string `json:"uuid"`
	Token       string `json:"token"`
	HouseNumber string `json:"house_number"`
	UserType    int    `json:"user_type"`
}
