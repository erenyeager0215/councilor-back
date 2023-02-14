package controllers

import (
	"fmt"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func session(c echo.Context) (sess models.Session, err error) {
	cookie, err := c.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		// cookieにある情報をDBにある情報と一致しているかチェック
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("セッションがありません")
		}
	}
	return sess, err
}

func StartMainServer() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	// Debug mode
	e.Debug = true

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/councilor/:id", getCouncilor)
	e.GET("/councilors", getCouncilors)

	e.GET("/category", getCategory)
	e.GET("/questions/:id", getQuestionsByCouncilorId)

	//user情報をDBへ登録
	e.POST("/register_user", registerUser)
	e.POST("/login", login)
	e.POST("/favorite/councilor", registerUsersFavoriteCouncilor)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
