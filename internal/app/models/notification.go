package models

type Notification struct {
	ID            string
	Result        string
	Reason        string
	Advise        string
	UploadCompany string
	UploadPerson  string
	UploadTime    string
	FoodChainID   string
	Detail        FoodChainNode
	Type          int
}
