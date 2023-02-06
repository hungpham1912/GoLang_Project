package userRepositories

import (
	"CRUD/src/database"
	CreateUserDto "CRUD/src/system/module/controller/users/dto"
	model "CRUD/src/system/module/model/users"
	"fmt"
	"strconv"
)

func Find() []model.User {
	users, err := database.ConnectDB.Query("SELECT * FROM users")
	if err != nil {
		panic((err.Error()))
	}

	array := []model.User{}

	for users.Next() {
		var user model.User
		err = users.Scan(&user.Id, &user.Name, &user.Email, &user.PassWord)
		if err != nil {
			panic((err.Error()))
		}

		array = append(array, user)
	}
	return array
}

func Save(newUser CreateUserDto.UserDto) model.User {
	queryStr := "INSERT INTO users (name,email,passWord) VALUES('" + newUser.Name + "','" + newUser.Email + "','" + newUser.PassWord + "')"

	user, err := database.ConnectDB.Query(queryStr)

	fmt.Println(newUser.PassWord)

	if err != nil {
		panic((err.Error()))
	}

	fmt.Println(user)

	var id int = 1

	tx, err := database.ConnectDB.Begin()

	if err != nil {
		panic((err.Error()))
	}

	{
		stmt, err := tx.Prepare(" SELECT MAX( id ) FROM users;")
		if err != nil {
			panic((err.Error()))
		}
		err = stmt.QueryRow().Scan(&id)

		fmt.Println(id)

		if err != nil {
			panic((err.Error()))
		}
	}

	{
		err := tx.Commit()

		if err != nil {
			panic((err.Error()))
		}
	}

	var newId string = strconv.Itoa(id)
	return FindOne(newId)

}

func FindOne(newId string) model.User {

	queryStr := "SELECT * FROM users WHERE id =" + newId
	users, err := database.ConnectDB.Query(queryStr)

	if err != nil {
		panic((err.Error()))
	}

	array := []model.User{}

	for users.Next() {
		var user model.User
		err = users.Scan(&user.Id, &user.Name, &user.Email, &user.PassWord)
		if err != nil {
			panic((err.Error()))
		}

		array = append(array, user)
	}

	return array[0]
}

func FindByEmail(email string) []model.User {
	queryStr := "SELECT * FROM users WHERE email =" + "'" + email + "'"
	users, err := database.ConnectDB.Query(queryStr)

	if err != nil {
		panic((err.Error()))
	}

	array := []model.User{}

	for users.Next() {
		var user model.User
		err = users.Scan(&user.Id, &user.Name, &user.Email, &user.PassWord)
		if err != nil {
			panic((err.Error()))
		}

		array = append(array, user)
	}

	return array
}
