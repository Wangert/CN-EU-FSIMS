package response

type ResAddOperator struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type ResTransportVehicle struct {
	TVNumber      string `json:"tv_number"`
	LicenseNumber string `json:"license_number"`
	Driver        string `json:"driver"`
	DriverPhone   string `json:"driver_phone"`
	State         uint   `json:"state"`
}

type ResHouse struct {
	HouseNumber string `json:"house_number"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	State       uint   `json:"state"`
	LegalPerson string `json:"legal_person"`
}
