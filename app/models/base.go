package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"myapp/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const(
	tabelNameUser = "users"
	tableNameCouncilor= "councilors"
	tableNameQuestion= "questions"
)

func init(){
	Db,err = sql.Open(config.Config.SQLDriver,config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nickname STRING NOT NULL UNIQUE,
		password STRING NOT NULL,
		created_at DATETIME,
		age STRING)`,tabelNameUser)

	_,err = Db.Exec(cmdU)
	if err != nil{
		log.Fatalln(err)
	}

	cmdC:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		commitee STRING,
		imagepath STRING,
		address STRING,
		tel_num STRING,
		birthday DATETIME,
		created_at DATETIME)`,tableNameCouncilor)

		_,err = Db.Exec(cmdC)
		
		if err != nil{
			log.Fatalln(err)
		}
	
		cmdQ := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			overview STRING,
			category STRING,
			content TEXT,
			answer TEXT,
			held_time STRING,
			councilor_id INTEGER,
			created_at DATETIME)`, tableNameQuestion)
	
		_,err:=Db.Exec(cmdQ)

		if err != nil{
			log.Fatalln(err)
		}
}

func Encrypt(plaintext string)(cryptext string){
	cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
	return cryptext
}