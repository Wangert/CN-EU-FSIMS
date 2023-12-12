package request

type ReqInPastureWarehouse struct {
	ProductPID  string `json:"product_pid" form:"product_pid"`
	ProductType string `json:"product_type" form:"product_type"`
	InOperator  string `json:"in_operator" form:"in_operator"`
}

type ReqSendToSlaughter struct {
	ProductNumber string `json:"product_number" form:"product_number"`
	OutOperator   string `json:"out_operator" form:"out_operator"`
	Destination   string `json:"destination" form:"destination"`
}

type SlaughterReceive struct {
	ProductNumber string `json:"product_number" form:"product_number"`
	ProductPID    string `json:"product_pid" form:"product_pid"`
	ProductType   string `json:"product_type" form:"product_type"`
	Operator      string `json:"operator" form:"operator"`
}

type ReqInSlaughterWarehouse struct {
	ProductPID  string `json:"product_pid" form:"product_pid"`
	ProductType string `json:"product_type" form:"product_type"`
	InOperator  string `json:"in_operator" form:"in_operator"`
}

type ReqSendToPack struct {
	ProductNumber string `json:"product_number" form:"product_number"`
	OutOperator   string `json:"out_operator" form:"out_operator"`
	Destination   string `json:"destination" form:"destination"`
}

type PackReceive struct {
	ProductNumber string `json:"product_number" form:"product_number"`
	ProductPID    string `json:"product_pid" form:"product_pid"`
	ProductType   string `json:"product_type" form:"product_type"`
	Operator      string `json:"operator" form:"operator"`
}

type ReqInPackWarehouse struct {
	ProductPID  string `json:"product_pid" form:"product_pid"`
	ProductType string `json:"product_type" form:"product_type"`
	InOperator  string `json:"in_operator" form:"in_operator"`
}

type ReqSendToTransport struct {
	ProductNumber string `json:"product_number"`
	OutOperator   string `json:"out_operator"`
	Destination   string `json:"destination"`
}
