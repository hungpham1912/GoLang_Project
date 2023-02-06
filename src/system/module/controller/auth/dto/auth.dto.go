package dto

type LoginDto struct {
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
}

type RegisterDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
}
