package models

type User struct {
	ID        int
	Username  string
	Password  string
	Company   string
	Usertype  int
	Telephone string
	WarnNum   int
	Login     bool
	Place     string
}
