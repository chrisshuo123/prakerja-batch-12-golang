package controllers

import (
	"net/http"
	"prakerja12/middlewares"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var userRegister models.User
	c.Bind(&userRegister)

	result, _ := bcrypt.GenerateFromPassword(
		[]byte(userRegister.Password),
        bcrypt.DefaultCost,
	)
	userRegister.Password = string(result)
	err := repositories.Register(&userRegister)
	
	if err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{
		    Message: err.Error(),
            Status: false,
            Data: nil,
		})
    }

	var userResponse models.UserResponse
	userResponse.Id = userRegister.Id
	userResponse.Name = userRegister.Name
	userResponse.Email = userRegister.Email
	userResponse.Token = middlewares.GenerateJWTToken(
		userResponse.Id,
		userResponse.Name,
	)

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil registered successfully",
        Status: true,
        Data: userResponse,
	})
}