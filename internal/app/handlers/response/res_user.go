package response

type ResUser struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Type    int    `json:"type"`
	Role    string `json:"role"`
	Status  int    `json:"status"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
}

type ResLogin struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}
