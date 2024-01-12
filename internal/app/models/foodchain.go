package models

type Foodchain struct {
	Checkcode            string `json:"checkcode"`
	CurrentProductNumber string `json:"current_product_number"`
	CurrentLastProcedure string `json:"current_last_procedure"`

	PasturePID   string `json:"pasture_pid"`
	SlaughterPID string `json:"slaughter_pid"`
	PackagePID   string `json:"package_pid"`
	ColdchainPID string `json:"coldchain_pid"`
}
