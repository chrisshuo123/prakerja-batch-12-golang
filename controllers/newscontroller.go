package controllers

import (
	"net/http"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
)

func CreateNewsController(c echo.Context) error {
	var newsRequest models.News
	c.Bind(&newsRequest)

	err := repositories.AddNews(&newsRequest)

	if err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{
		    Message: err.Error(),
            Status: false,
            Data: nil,
		})
    }
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
        Status: true,
        Data: newsRequest,
	})
} 

func GetNewsController(c echo.Context) error {
	var news []models.News
	
	err := repositories.GetNews(&news)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
            Status: false,
            Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
        Status: true,
        Data: news,
	})
}