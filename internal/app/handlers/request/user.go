package request

type ReqUser struct {
	Name            string `json:"name" form:"name"`
	Account         string `json:"account" form:"account"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Type            int    `json:"type" form:"type"`
	Role            string `json:"role" form:"role"`
	Company         string `json:"company" form:"company"`
	Phone           string `json:"phone" form:"phone"`
}

type ReqLogin struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	Type     int    `json:"type" form:"type"`
}

type ReqAddUser struct {
	Name            string `json:"name" form:"name"`
	Account         string `json:"account" form:"account"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Type            int    `json:"type" form:"type"`
	Role            string `json:"role" form:"role"`
	Company         string `json:"company" form:"company"`
	Phone           string `json:"phone" form:"phone"`
}

type ReqUpdateUser struct {
	Name            string `json:"name" form:"name"`
	Account         string `json:"account" form:"account"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Type            int    `json:"type" form:"type"`
	Role            string `json:"role" form:"role"`
	Company         string `json:"company" form:"company"`
	Phone           string `json:"phone" form:"phone"`
}
