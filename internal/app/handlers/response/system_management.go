package response

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/coldchain"
	"CN-EU-FSIMS/internal/app/models/sell"
)

type ResAddOperator struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type ResTransportVehicles struct {
	TVs   []coldchain.TransportVehicleInfo `json:"tvs"`
	Count int64                            `json:"count"`
}

type ResHouses struct {
	Houses []models.House `json:"houses"`
	Count  int64          `json:"count"`
}

type ResMalls struct {
	Malls []sell.MallInfo `json:"malls"`
	Count int64           `json:"count"`
}
