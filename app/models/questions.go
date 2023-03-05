package models

import (
	"log"
	"time"
)

type Question struct {
	Id           int       `json:"id"`
	Overview     string    `json:"overview"`
	Category     string    `json:"category"`
	Content      string    `json:"content"`
	Answer       string    `json:"answer"`
	Held_time    string    `json:"held_time"`
	Councilor_id int       `json:"councilor_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type Category struct {
	Category_id   int       `json:"id"`
	Category_name string    `json:"category_name"`
	CreatedAt     time.Time `json:"created_at"`
}

func GetQuestions(id int) (questions []Question, err error) {
	cmd := "SELECT overview,category,content,answer,held_time,councilor_id from questions WHERE councilor_id = ?"
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var que Question
		err = rows.Scan(
			&que.Overview,
			&que.Category,
			&que.Content,
			&que.Answer,
			&que.Held_time,
			&que.Councilor_id,
		)
		if err != nil {
			log.Fatalln(err)
		}
		questions = append(questions, que)
	}
	return questions, err
}

func GetQuestionsByCategory(category string) (questions []Question, err error) {
	cmd := "SELECT overview,category,content,answer,held_time,councilor_id from questions WHERE category = ?"
	rows, err := Db.Query(cmd, category)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var que Question
		err = rows.Scan(
			&que.Overview,
			&que.Category,
			&que.Content,
			&que.Answer,
			&que.Held_time,
			&que.Councilor_id,
		)
		if err != nil {
			log.Fatalln(err)
		}
		questions = append(questions, que)
	}
	return questions, err
}

func GetQuestionList() (questionList []Question, err error) {
	cmd := `SELECT overview,category,content,answer,held_time,councilor_id from questions`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var que Question
		err = rows.Scan(
			&que.Overview,
			&que.Category,
			&que.Content,
			&que.Answer,
			&que.Held_time,
			&que.Councilor_id,
		)
		if err != nil {
			log.Fatalln(err)
		}
		questionList = append(questionList, que)
	}
	return questionList, err
}

func GetCategory() (categories []Category, err error) {
	cmd := `SELECT category.category_id,questions.category FROM questions JOIN category ON questions.category = category.category_name GROUP BY category.category_id`

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.Category_id,
			&category.Category_name,
		)
		if err != nil {
			log.Fatalln(err)
		}
		categories = append(categories, category)
	}
	return categories, err
}

func GetQuestionsByCategory2(id int) (questions []Question, err error) {
	cmd := `SELECT overview,category,content,answer,held_time,councilor_id FROM questions JOIN category ON questions.category = category.category_name WHERE category.category_id = ?`
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var question Question
		err = rows.Scan(
			&question.Overview,
			&question.Category,
			&question.Content,
			&question.Answer,
			&question.Held_time,
			&question.Councilor_id,
		)
		if err != nil {
			log.Fatalln(err)
		}
		questions = append(questions, question)
	}
	return questions, err
}
