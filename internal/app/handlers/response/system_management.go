package response

import "CN-EU-FSIMS/internal/app/models"

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

type ResHouses struct {
	Houses []models.House `json:"houses"`
	Count  int64          `json:"count"`
}
