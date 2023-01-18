package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"myapp/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const(
	tabelNameUser = "users"
	tableNameCouncilor= "councilors"
	tableNameQuestion= "questions"
	tableNameTest= "test_table"
	tableNameSession= "session"
	tableNameFavorite= "favorite"
)

func init(){
	Db,err = sql.Open(config.Config.SQLDriver,config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING,
		nickname STRING NOT NULL UNIQUE,
		password STRING NOT NULL,
		created_at DATETIME
		)`,tabelNameUser)

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

		cmdT:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name STRING,
			commitee STRING,
			image STRING,
			address STRING,
			contact STRING,
			birthday DATETIME,
			url STRING,
			created_at DATETIME)`,tableNameTest)
	
		_,err = Db.Exec(cmdT)
		if err != nil{
			log.Fatalln(err)
		}

		cmdS:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			nickname STRING,
			user_id INTEGER,
			created_at DATETIME)`,tableNameSession)
	
		_,err = Db.Exec(cmdS)
		if err != nil{
			log.Fatalln(err)
		}

		cmdF:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			councilor_id INTEGER,
			category STRING,
			created_at DATETIME)`,tableNameFavorite)
	
		_,err = Db.Exec(cmdF)
		if err != nil{
			log.Fatalln(err)
		}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string)(cryptext string){
	cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
	return cryptext
}