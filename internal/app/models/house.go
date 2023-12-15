package models

type House struct {
	HouseNumber string `json:"house_number"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	State       uint   `json:"state"`
	LegalPerson string `json:"legal_person"`
}
