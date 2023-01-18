package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
)

func registerUsersFavoriteCouncilor(c echo.Context)error{
	uf:= new(models.UserFav)
	if err:= c.Bind(uf); err !=nil{
		log.Println(uf.User_id)
		log.Println(uf.Councilor_id)
		log.Fatalln(err)
	}
	user_fav,err:=models.GetFavoriteCouncilorByUserId(uf.User_id)
	if err !=nil{
		log.Println(err)
		// もし、ユーザが支持する議員がいない場合新規で支持者を登録する
		ufFromPost,err:=uf.PostFavoriteCouncilor()
		if err!=nil{
			log.Println(err)
		}
		return c.JSON(http.StatusCreated,ufFromPost)	
	}
	if user_fav.Councilor_id != uf.Councilor_id{
		// もし、ユーザが支持する議員が以前の議員から変更があった場合、議員情報を更新する。
		ufFromPut,err:=uf.PutFavoriteCouncilor()
		if err!=nil{
			log.Println(err)
		}
		return c.JSON(http.StatusCreated,ufFromPut)
	} 
	return c.JSON(http.StatusCreated,user_fav)
}