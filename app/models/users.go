package models

import (
	"log"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	NickName  string    `json:"nickname"`
	PassWord  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	Gender    string    `json:"gender"`
	Home      string    `json:"home"`
	CreatedAt time.Time `json:"created_at"`
}

type UserData struct {
	GroupAge    string `json:"group_age"`
	NumOfPeople int    `json:"num_of_people"`
}

type ResponseUser struct {
	UserID   int     `json:"user_id"`
	Favorite UserFav `json:"favorite"`
}

type Session struct {
	Id        int
	Uuid      string
	NickName  string
	UserId    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users(
		uuid,
		nickname,
		password,
		birthday,
		created_at,
		gender,
		home
	) VALUES(?,?,?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.NickName,
		Encrypt(u.PassWord),
		u.Birthday,
		time.Now(),
		u.Gender,
		u.Home,
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// ニックネームに紐づくユーザ情報を取得
func GetUser(u *User) (user User, err error) {
	cmd := "SELECT id,nickname,password from users WHERE nickname = ?"
	err = Db.QueryRow(cmd, u.NickName).Scan(
		&user.ID,
		&user.NickName,
		&user.PassWord,
	)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

func GetUserData() (userData []UserData, err error) {
	// 年代別ユーザ数を取得
	cmd := "SELECT CASE WHEN age <= 19 THEN '19s' WHEN age BETWEEN 20 AND 29 THEN '20s' WHEN age BETWEEN 30 AND 39 THEN '30s' WHEN age BETWEEN 40 AND 49 THEN '40s' WHEN age BETWEEN 50 AND 59 THEN '50s' WHEN age BETWEEN 60 AND 69 THEN '60s' WHEN age BETWEEN 70 AND 79 THEN '70s' ELSE '80~'  END AS GroupAge,COUNT(*) AS NumOfPeople FROM (SELECT id,TIMESTAMPDIFF(YEAR, `birthday`, CURDATE()) AS age FROM users)AS age_table GROUP BY GroupAge ORDER BY 1 ASC"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var ud UserData
		err = rows.Scan(
			&ud.GroupAge,
			&ud.NumOfPeople,
		)
		if err != nil {
			log.Fatalln(err)
		}
		userData = append(userData, ud)
	}
	return userData, err
}

func (user *User) CreateSession() (sesson Session, err error) {
	session := Session{}
	cmd1 := `INSERT INTO session(
		uuid,
		nickname,
		user_id,
		created_at
	)values(?,?,?,?)`

	_, err = Db.Exec(cmd1,
		createUUID(),
		user.NickName,
		user.ID,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `SELECT id,uuid,nickname,user_id,created_at from session where nickname=? and user_id = ?`
	err = Db.QueryRow(cmd2, user.NickName, user.ID).Scan(
		&sesson.Id,
		&sesson.Uuid,
		&sesson.NickName,
		&sesson.UserId,
		&sesson.CreatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT id,uuid,nickname,user_id,created_at from session where uuid = ?`

	err = Db.QueryRow(cmd, sess.Uuid).Scan(
		&sess.Id,
		&sess.Uuid,
		&sess.NickName,
		&sess.UserId,
		&sess.CreatedAt,
	)
	if err != nil {
		valid = false
		return
	}
	if sess.Id != 0 {
		valid = true
	}
	return valid, err
}
