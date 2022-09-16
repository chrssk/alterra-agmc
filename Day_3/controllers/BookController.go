package Controller

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"Day2Assignment/models"
)


func GetAllBookController(c echo.Context) error {

	books := []string{"5 Sekawan", "Hardy Boys", "St Clare", "Mallory Towers" }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

func GetBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{Id: id, Judul: "Mallory Towers", Penulis: "Enid Blyton"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func CreateNewBookController(c echo.Context) error{
	Judul := c.FormValue("Judul")
	Penulis := c.FormValue("Penulis")

	var book models.Book
	book.Judul = Judul
	book.Penulis = Penulis
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create book",
		"book": book,
	})

}

func UpdateBookController(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	Judul := c.FormValue("Judul")
	Penulis := c.FormValue("Penulis")

	var book models.Book
	book.Id = id
	book.Judul = Judul
	book.Penulis = Penulis
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book",
		"book": book,
	})
}

func DeleteBookController(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	IDstr := strconv.Itoa(id)
	messages := "success deleted book with id :" + IDstr

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})

}