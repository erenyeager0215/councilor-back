package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	NickName  string `json:"nickname" form:"nickname" query:"nickname"`
	PassWord  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User)CreateUser()(err error){
	cmd:= `INSERT INTO users(
		nickname,
		password,
		created_at
	) VALUES(?,?,?)`

	_,err= Db.Exec(cmd,
		u.NickName,
		Encrypt(u.PassWord),
		time.Now(),
	)
	if err != nil{
		log.Fatal(err)
	}
	return err
}




func GetUser(u *User)(user User,err error){
	cmd:= "SELECT nickname,password from users WHERE nickname = ?"
	err= Db.QueryRow(cmd,u.NickName).Scan(			
			&user.NickName,
			&user.PassWord,
		)
		if err != nil {
			log.Fatal(err)
		}
		return user,err
}