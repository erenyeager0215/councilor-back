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
	questions, err := models.GetQuestions(i)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, questions)
}

func getQuestions(c echo.Context) error {
	questionList, err := models.GetQuestionList()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, questionList)
}

func getCategory(c echo.Context) error {
	categories, err := models.GetCategory()
	log.Println(categories)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, categories)
}

func getQuestionsByCategory(c echo.Context) error {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		log.Println(err)
	}
	questions, err := models.GetQuestionsByCategory2(id)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusCreated, questions)
}
