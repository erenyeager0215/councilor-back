package models

import (
	"log"
	"time"
)

type UserFav struct {
	Id           int       `json:"id"`
	User_id      int       `json:"user_id,omitempty"`
	Councilor_id int       `json:"councilor_id"`
	Category_id  int       `json:"category_id"`
	Created_at   time.Time `json:"-"`
}

func GetFavoriteByUserId(id int) (user_fav UserFav, err error) {
	cmd := `SELECT id,user_id,councilor_id,category_id FROM favorite WHERE user_id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user_fav.Id,
		&user_fav.User_id,
		&user_fav.Councilor_id,
		&user_fav.Category_id,
	)
	if err != nil {
		e := RaiseError()
		log.Print(e)
		return
	}
	return user_fav, err
}

func (user_fav *UserFav) PostFavoriteCouncilor() (UserFav, error) {
	cmd1 := `INSERT INTO favorite(user_id,councilor_id) VALUES(?,?)`
	_, err = Db.Exec(cmd1, user_fav.User_id, user_fav.Councilor_id)
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `SELECT id,user_id,councilor_id,category_id FROM favorite WHERE user_id = ?`
	err = Db.QueryRow(cmd2, user_fav.User_id).Scan(
		&user_fav.Id,
		&user_fav.User_id,
		&user_fav.Councilor_id,
		&user_fav.Category_id,
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("新規に支持する議員を登録しました")
	return *user_fav, err
}

func (user_fav *UserFav) PutFavoriteCouncilor() (UserFav, error) {
	cmd1 := `UPDATE favorite SET councilor_id = ? WHERE user_id = ?`
	_, err := Db.Exec(cmd1, user_fav.Councilor_id, user_fav.User_id)
	if err != nil {
		log.Fatalln(err)
	}
	cmd2 := `SELECT id,user_id,councilor_id,category_id FROM favorite WHERE user_id = ?`
	err = Db.QueryRow(cmd2, user_fav.User_id).Scan(
		&user_fav.Id,
		&user_fav.User_id,
		&user_fav.Councilor_id,
		&user_fav.Category_id,
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("支持する議員を更新しました")
	return *user_fav, err
}

func (user_fav *UserFav) PutFavoriteCategory() (UserFav, error) {
	cmd1 := `UPDATE favorite SET category_id = ? WHERE user_id = ?`
	_, err := Db.Exec(cmd1, user_fav.Category_id, user_fav.User_id)
	if err != nil {
		log.Fatalln(err)
	}
	cmd2 := `SELECT id,user_id,councilor_id,category_id FROM favorite WHERE user_id = ?`
	err = Db.QueryRow(cmd2, user_fav.User_id).Scan(
		&user_fav.Id,
		&user_fav.User_id,
		&user_fav.Councilor_id,
		&user_fav.Category_id,
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("カテゴリを更新しました")
	return *user_fav, err
}
