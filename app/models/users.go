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
	CreatedAt time.Time `json:"created_at"`
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
		created_at
	) VALUES(?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.NickName,
		Encrypt(u.PassWord),
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

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
