package model

type User struct {
	CommonModel
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
