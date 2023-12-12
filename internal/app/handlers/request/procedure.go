package request

type ReqCreateProcedure struct {
	Type     uint   `json:"type" form:"type"`
	Operator string `json:"operator" form:"operator"`
	PrePID   string `json:"pre_pid" form:"pre_pid"`
}

type CommitPastureProcedure struct {
	PID             string  `json:"pid" form:"pid"`
	TotalBacteria   uint    `json:"total_bacteria"`   // 总细菌数
	AmmoniaGas      float32 `json:"ammonia_gas"`      // 氨气
	HydrogenSulfide float32 `json:"hydrogen_sulfide"` // 硫化氢
	CarbonDioxide   float32 `json:"carbon_dioxide"`   // 二氧化碳
	PM10            float32 `json:"pm_10"`            // PM10
	TSP             float32 `json:"tsp"`              // TSP
	Stench          uint    `json:"stench"`           // 恶臭稀释程度
}

type CommitSlaughterProcedure struct {
	PID                   string `json:"pid" form:"pid"`
	EnvirTemperature      string `json:"envir_temperature"`
	EnvirLighting         string `json:"envir_lighting"`
	ShockVoltage          string `json:"shock_voltage"`
	BleedingTime          string `json:"bleeding_time"`
	EleShockTime          string `json:"ele_shock_time"`
	KnifeDisinfectionTime string `json:"knife_disinfection_time"`
}

type CommitPackProcedure struct {
	PID             string `json:"pid" form:"pid"`
	PackType        string `json:"pack_type" form:"pack_type"`
	PackTemperature string `json:"pack_temperature" form:"pack_temperature"`
}

type TransportStart struct {
	Type          uint   `json:"type" form:"type"`
	Operator      string `json:"operator" form:"operator"`
	CarNumber     string `json:"car_number" form:"car_number"`
	PrePID        string `json:"pre_pid" form:"pre_pid"`
	ProductNumber string `json:"product_number" form:"product_number"`
	Temperature   string `json:"temperature" form:"temperature"`
	Source        string `json:"source" form:"source"`
	Destination   string `json:"destination" form:"destination"`
	Humidity      string `json:"humidity" form:"humidity"`
	LoadingTime   string `json:"loading_time" form:"loading_time"`
}

type TransportEnd struct {
	PID           string `json:"pid" form:"pid"`
	UnloadingTime string `json:"unloading_time" form:"unloading_time"`
}
