package authController

import (
	"CRUD/src/system/module/controller/auth/dto"
	service "CRUD/src/system/module/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Login(context *gin.Context) {
// 	var newUser dto.UserDto
// 	result := service.GetAllUsers()
// 	context.IndentedJSON(http.StatusCreated, result)
// }

func Register(context *gin.Context) {
	var registerUser dto.RegisterDto
	context.BindJSON(&registerUser)
	result := service.Register(registerUser)
	context.IndentedJSON(http.StatusCreated, result)
}
