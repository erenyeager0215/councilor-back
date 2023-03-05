package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
)

func registerUsersFavoriteCouncilor(c echo.Context) error {
	uf := new(models.UserFav)
	if err := c.Bind(uf); err != nil {
		log.Fatalln(err)
	}

	// レスポンス用の構造体へ格納
	responseUser := &models.ResponseUser{}
	user_fav, err := models.GetFavoriteByUserId(uf.User_id)

	// もし、ユーザが支持する議員がいない場合新規で支持者を登録する
	if err != nil {
		log.Println(err)
		user_newfav, err := uf.PostFavoriteCouncilor()
		if err != nil {
			log.Println(err)
		}
		responseUser.UserID = user_newfav.User_id
		responseUser.Favorite = user_newfav
		return c.JSON(http.StatusCreated, responseUser)
	}

	// もし、ユーザが支持する議員が以前の議員から変更があった場合、議員情報を更新する。
	if user_fav.Councilor_id != uf.Councilor_id {
		user_newfav, err := uf.PutFavoriteCouncilor()
		if err != nil {
			log.Println(err)
		}
		responseUser.UserID = user_newfav.User_id
		responseUser.Favorite = user_newfav
		return c.JSON(http.StatusCreated, responseUser)
	}

	responseUser.UserID = user_fav.User_id
	responseUser.Favorite = user_fav
	return c.JSON(http.StatusCreated, responseUser)
}

func registerUsersFavoriteCategory(c echo.Context) error {
		userFav:=new(models.UserFav)
		if err:= c.Bind(userFav); err != nil{
			log.Fatalln(err)
		}

		responseUserFav := &models.ResponseUser{}
		userFavCategory,err := 

}
