package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		log.Fatal(err)
	}

	user, err := models.GetUser(u)
	if err != nil {
		log.Fatalln(err)
	}
	
	if models.Encrypt(u.PassWord) == user.PassWord {
		return c.JSON(http.StatusCreated, "OK")
	} else {
		return c.JSON(http.StatusCreated, "NotFound")
	}
}

func registerUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	log.Print(u)

	err := u.CreateUser()
	if err != nil {
		log.Fatalln(err)
	}
	return c.JSON(http.StatusCreated, "OK")
}

