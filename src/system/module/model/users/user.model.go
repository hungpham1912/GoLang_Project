package userModel

import "CRUD/src/system/module/model"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"passWord"`
}

type ResponseGetAllUsers struct {
	model.Response
	Data []User `json:"data"`
}

type ResponseGetAnUsers struct {
	model.Response
	Data User `json:"data"`
}

type ResponseRegister struct {
	model.Response
	Data        User   `json:"data"`
	AccessToken string `json:"accessToken"`
}
