package models

type UserLoginModel struct {
	User     string `json:"username"`
	Password string `json:"password"`
}
