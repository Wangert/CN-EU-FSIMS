package request

type ReqAddOperator struct {
	Name        string `json:"name" form:"name"`
	Account     string `json:"account" form:"account"`
	Type        int    `json:"type" form:"type"`
	Role        string `json:"role" form:"role"`
	Company     string `json:"company" form:"company"`
	Phone       string `json:"phone" form:"phone"`
	HouseNumber string `json:"house_number" form:"house_number"`
}

type ReqSearchUser struct {
	Name        string `json:"name" form:"name"`
	Role        string `json:"role"`
	Company     string `json:"company" form:"company"`
	HouseNumber string `json:"house_number" form:"house_number"`
}

type ReqAddPasture struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	State       uint   `json:"state" form:"state"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqSearchPasture struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqAddSlaughterHouse struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	State       uint   `json:"state" form:"state"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqSearchSlaughterHouse struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqAddPackHouse struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	State       uint   `json:"state" form:"state"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqSearchPackHouse struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	LegalPerson string `json:"legal_person" form:"legal_person"`
}

type ReqAddVehicle struct {
	LicenseNumber string `json:"license_number" form:"license_number"`
	Driver        string `json:"driver" form:"driver"`
	DriverPhone   string `json:"driver_phone" form:"driver_phone"`
}

type ReqSearchVehicle struct {
	LicenseNumber string `json:"license_number" form:"license_number"`
	Driver        string `json:"driver" form:"driver"`
	DriverPhone   string `json:"driver_phone" form:"driver_phone"`
}
