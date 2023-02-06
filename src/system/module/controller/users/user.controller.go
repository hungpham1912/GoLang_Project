package userController

import (
	CreateUserDto "CRUD/src/system/module/controller/users/dto"
	service "CRUD/src/system/module/service/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	result := service.GetAllUsers()
	context.IndentedJSON(http.StatusCreated, result)
}

func AddOneUser(context *gin.Context) {
	var newUser CreateUserDto.UserDto
	context.BindJSON(&newUser)
	user := service.AddOneUser(newUser)
	context.IndentedJSON(http.StatusOK, user)
}

func GetOneUser(context *gin.Context) {
	userId := context.Param("id")
	user := service.GetOneUser(userId)
	context.IndentedJSON(http.StatusOK, user)

}
