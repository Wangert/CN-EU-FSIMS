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
