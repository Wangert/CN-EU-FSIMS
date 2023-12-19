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
