package userService

import (
	userRepositories "CRUD/src/repositories"
	CreateUserDto "CRUD/src/system/module/controller/users/dto"
	model "CRUD/src/system/module/model/users"
)

func GetAllUsers() model.ResponseGetAllUsers {
	listUser := userRepositories.Find()

	var result model.ResponseGetAllUsers
	result.Status = true
	result.Message = "OK"
	result.Data = listUser

	return result
}

func AddOneUser(newUser CreateUserDto.UserDto) model.User {
	ds := userRepositories.Save(newUser)
	return ds
}

func GetOneUser(userId string) model.ResponseGetAnUsers {
	user := userRepositories.FindOne(userId)
	var result model.ResponseGetAnUsers
	result.Status = true
	result.Message = "OK"
	result.Data = user
	return result
}
