package main
import (
	//"net/http"
	//"github.com/labstack/echo/v4"
	//"strconv"
	"Day2Assignment/config"
	"Day2Assignment/routes"
	m "Day2Assignment/middlewares"
)

func main(){
	e := route.New()

	config.InitDB()
	config.InitialMigration()

	m.LogMiddleware(e)
	e.Start(":8080")
}



