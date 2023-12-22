package request

type ReqNewPackageBatch struct {
	HouseNumber   string `json:"house_number" form:"house_number"`
	Worker        string `json:"worker" form:"worker"`
	PrePID        string `json:"pre_pid" form:"pre_pid"`
	ProductNumber string `json:"product_number" form:"product_number"`
}

type ReqEndPackage struct {
	BatchNumber string `json:"batch_number" form:"batch_number"`
	Worker      string `json:"worker" form:"worker"`
	HouseNumber string `json:"house_number" form:"house_number"`

	PackType        string `json:"pack_type"`
	PackTemperature string `json:"pack_temperature"`
}

type ReqNewPackageProduct struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	Type        int     `json:"type" form:"type"`
	ShelfLife   string  `json:"shelf_life" form:"shelf_life"`
	PackMethod  int     `json:"pack_method" form:"pack_method"`
	Weight      float64 `json:"weight" form:"weight"`
}

type ReqSendToTransportVehicle struct {
	ProductNumber string `json:"product_number" form:"product_number"`
	Operator      string `json:"operator" form:"operator"`
	VehicleNumber string `json:"vehicle_number" form:"vehicle_number"`
}
