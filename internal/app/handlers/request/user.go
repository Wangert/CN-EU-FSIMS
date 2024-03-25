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
	Name    string `json:"name" form:"name"`
	Account string `json:"account" form:"account"`
	Type    int    `json:"type" form:"type"`
	Role    string `json:"role" form:"role"`
	Company string `json:"company" form:"company"`
	Phone   string `json:"phone" form:"phone"`
}

// reset password or delete user by admin
type ReqAccount struct {
	Account string `json:"account" form:"account"`
	Type    int    `json:"type" form:"type"`
}

type ReqTrashPerDay struct {
	TimeStamp int64 `json:"time_stamp" form:"time_stamp"`
}

type ReqTrashFifteenDays struct {
	StartTimeStamp int64 `json:"start_time_stamp" form:"start_time_stamp"`
	EndTimeStamp   int64 `json:"end_time_stamp" form:"end_time_stamp"`
}

type ReqSlPrediction struct {
	Id          uint    `json:"id" form:"id"`
	Temperature float64 `json:"temperature" form:"temperature"`
}

type ReqCleanImage struct {
	Id uint `json:"id" form:"id"`
}
