package Controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"Day2Assignment/models"
	"Day2Assignment/lib"
)

func LoginUsersController(c echo.Context) error {
	users, e := database.GetAllUser()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users" : users,
	})
}

func GetAllUsersController(c echo.Context) error {
	users, e := database.GetAllUser()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users" : users,
	})
}

func GetUserByIdController(c echo.Context) error {
	user, e := database.GetUserById(c.Param("id"))

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user" : user,
	})
}

func CreateNewUserController(c echo.Context) error{
	Name := c.FormValue("Name")
	Email := c.FormValue("Email")
	Password := c.FormValue("Password")

	var User models.User
	User.Name = Name
	User.Email = Email
	User.Password = Password
	e := database.CreateUser(User)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})

}

func UpdateUserController(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	Name := c.FormValue("Name")
	Email := c.FormValue("Email")

	var User models.User
	User.Id = id
	User.Name = Name
	User.Email = Email
	e := database.UpdateUser(c.Param("id"), User)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": e,
	})
}

func DeleteUserController(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	IDstr := strconv.Itoa(id)
	database.DeleteUser(c.Param("id"))
	messages := "success deleted user with id :" + IDstr

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})

}