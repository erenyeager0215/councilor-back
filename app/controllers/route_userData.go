package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
)

func getUserData(c echo.Context) error {
	var userData []models.UserData
	userData, err := models.GetUserData()
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusCreated, userData)
}
