package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)


func StartMainServer() {
	e := echo.New()

	// Debug mode
	e.Debug = true

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/councilor",getCouncilor)
	e.GET("/councilors",getCouncilors)

	//user情報をDBへ登録
	e.POST("/register_user",registerUser)

	e.POST("/login",login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}


