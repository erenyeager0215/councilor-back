package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


func StartMainServer() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	  }))

	// Debug mode
	e.Debug = true

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/councilor/:id",getCouncilor)
	e.GET("/councilors",getCouncilors)

	e.GET("/category",getCategory)
	e.GET("/questions/:id",getQuestionsByCouncilorId)

	//user情報をDBへ登録
	e.POST("/register_user",registerUser)

	e.POST("/login",login)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}


