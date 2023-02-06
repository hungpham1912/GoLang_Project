package main

import (
	"CRUD/src/database"
	authController "CRUD/src/system/module/controller/auth"
	userController "CRUD/src/system/module/controller/users"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()
	/*Users Router*/
	router.GET("/users", userController.GetAllUsers)
	router.POST("/users", userController.AddOneUser)
	router.GET("/users/:id", userController.GetOneUser)
	/*Auth Router*/
	router.POST("/auth/register", authController.Register)

	/*Listening Server*/
	router.Run("localhost:9090")

}
