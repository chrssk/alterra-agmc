package main
import (
	//"net/http"
	//"github.com/labstack/echo/v4"
	//"strconv"
	"Day2Assignment/config"
	"Day2Assignment/routes"
)

func main(){
	config.InitDB()
	config.InitialMigration()
	e := route.New()
	e.Start(":8080")
}



