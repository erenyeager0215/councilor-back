package controllers

import (
	"log"
	"myapp/app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getQuestionsByCouncilorId(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Print("文字列をキャッチ")
		questions, err := models.GetQuestionsByCategory(id)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusCreated, questions)
	}
	log.Print("数字をキャッチ")
	questions, err := models.GetQuestions(i)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, questions)
}


func getQuestions(c echo.Context)error{
	questionList,err:= models.GetQuestionList()
	if err != nil{
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated,questionList)
}

func getCategory(c echo.Context)error{
	categories,err:=models.GetQuestionsCategory()
	if err != nil{
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated,categories)
}