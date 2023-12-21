package request

type ReqAddCow struct {
	Age         int     `json:"age" form:"age"`
	Weight      float64 `json:"weight" form:"weight"`
	HouseNumber string  `json:"house_number" form:"house_number"`
}

type ReqNewFeedingBatch struct {
	HouseNumber string   `json:"house_number" form:"house_number"`
	Worker      string   `json:"worker" form:"worker"`
	PrePID      string   `json:"pre_pid" form:"pre_pid"`
	CowNumbers  []string `json:"cow_numbers" form:"cow_numbers"`
}

type ReqEndFeeding struct {
	BatchNumber string  `json:"batch_number" form:"batch_number"`
	Worker      string  `json:"worker" form:"worker"`
	HouseNumber string  `json:"house_number" form:"house_number"`
	PM10        float32 `json:"pm_10" form:"pm_10"`   // PM10
	TSP         float32 `json:"tsp" form:"tsp"`       // TSP
	Stench      uint    `json:"stench" form:"stench"` // 恶臭稀释程度
}

type ReqSendToSlaughter struct {
	CowNumber            string `json:"cow_number" form:"cow_number"`
	Operator             string `json:"operator" form:"operator"`
	SlaughterHouseNumber string `json:"slaughter_house_number" form:"slaughter_house_number"`
}
