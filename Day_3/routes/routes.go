package route

import (
	//"net/http"
	"github.com/labstack/echo/v4"
	//"strconv"
	//"Day2Assignment/config"
	"Day2Assignment/controllers"
)

func New() *echo.Echo{
	e := echo.New()

	e.GET("/books", Controller.GetAllBookController)
	e.GET("/books/:id", Controller.GetBookByIdController)
	e.POST("/books", Controller.CreateNewBookController)
	e.PUT("/books/:id", Controller.UpdateBookController)
	e.DELETE("/books/:id", Controller.DeleteBookController)

	e.GET("/users", Controller.GetAllUsersController)
	e.GET("/users/:id", Controller.GetUserByIdController)
	e.POST("/users", Controller.CreateNewUserController)
	e.PUT("/users:id", Controller.UpdateUserController)
	e.DELETE("/users/:id", Controller.DeleteUserController)

	return e
}