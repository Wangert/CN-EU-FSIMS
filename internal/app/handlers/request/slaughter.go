package request

type ReqNewSlaughterBatch struct {
	HouseNumber string `json:"house_number" form:"house_number"`
	Worker      string `json:"worker" form:"worker"`
	PrePID      string `json:"pre_pid" form:"pre_pid"`
	CowNumber   string `json:"cow_number" form:"cow_number"`
}

type ReqEndSlaughter struct {
	BatchNumber string `json:"batch_number" form:"batch_number"`
	Worker      string `json:"worker" form:"worker"`
	HouseNumber string `json:"house_number" form:"house_number"`

	EnvirTemperature      string `json:"envir_temperature"`
	EnvirLighting         string `json:"envir_lighting"`
	ShockVoltage          string `json:"shock_voltage"`
	BleedingTime          string `json:"bleeding_time"`
	EleShockTime          string `json:"ele_shock_time"`
	KnifeDisinfectionTime string `json:"knife_disinfection_time"`
}

type ReqNewSlaughterProduct struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	Type        int     `json:"type" form:"type"`
	Weight      float64 `json:"weight" form:"weight"`
}

type ReqSendToPackage struct {
	ProductNumber      string `json:"product_number" form:"product_number"`
	Operator           string `json:"operator" form:"operator"`
	PackageHouseNumber string `json:"package_house_number" form:"package_house_number"`
}
