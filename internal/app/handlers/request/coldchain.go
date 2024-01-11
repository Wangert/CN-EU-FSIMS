package request

type ReqStartTransport struct {
	BatchNumber string `json:"batch_number" form:"batch_number"`
	Operator    string `json:"operator" form:"operator"`
}

type ReqEndTransport struct {
	BatchNumber string `json:"batch_number" form:"batch_number"`
	Worker      string `json:"worker" form:"worker"`

	Temperature string `json:"temperature" form:"temperature"`
	Source      string `json:"source" form:"source"`
	Destination string `json:"destination" form:"destination"`
	Humidity    string `json:"humidity" form:"humidity"`
}
