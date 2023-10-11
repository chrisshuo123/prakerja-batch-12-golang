package main

import (
	"prakerja12/configs"
	"prakerja12/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	configs.LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}






