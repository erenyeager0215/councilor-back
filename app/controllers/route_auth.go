package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"
	"time"

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
		sess,err:=user.CreateSession()
		if err != nil{
			log.Println(err)
		}

		cookie := new(http.Cookie)
		cookie.Name = "_cookie"
		cookie.Value = sess.Uuid
		cookie.Expires = time.Now().Add(time.Hour)
		cookie.HttpOnly = true
		c.SetCookie(cookie)	

		// ユーザに紐づくfavorite情報を取得
		user_fav,err:=models.GetFavoriteCouncilorByUserId(user.ID)
		if err != nil{
			log.Println(err)
		}

		// レスポンス用の構造体へ格納
		responseUser:= &models.ResponseUser{}
		responseUser.UserID=user.ID
		responseUser.Favorite=user_fav

		// ログインしたらフロントエンドへユーザIDを送る
		return c.JSON(http.StatusCreated, responseUser)
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


