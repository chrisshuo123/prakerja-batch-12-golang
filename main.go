package main

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Email string
}

func main(){
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetDetailUsersController)
	e.Start(":8000")
}

func GetUsersController(c echo.Context) error {
	country := c.QueryParam("country")

	var users = []User{{country, "alta@gmail.com"}}
	return c.JSON(200, users)
}

func GetDetailUsersController(c echo.Context) error {
	id := c.Param("id")
	var user = User{id, "alta@gmail.com"}
	return c.JSON(200, user)
}
