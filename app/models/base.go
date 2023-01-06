package models

import (
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
		created_at DATETIME)`,tabelNameUser)

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
}