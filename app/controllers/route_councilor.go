package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getCouncilor(c echo.Context) error{
	id := c.Param("id")
	i,_:= strconv.Atoi(id)
	councilor, err := models.GetCouncilor(i)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, councilor)
}


func getCouncilors(c echo.Context) error {
	var councilors []models.Councilor
	councilors, err := models.GetCouncilorList()
	if err != nil {
		log.Fatalln(err)
	}
	return c.JSON(http.StatusCreated, councilors)
}