package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"myapp/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var Db *sql.DB

var err error

const (
	tabelNameUser      = "users"
	tableNameCouncilor = "councilors"
	tableNameQuestion  = "questions"
	tableNameSession   = "session"
	tableNameFavorite  = "favorite"
)

func init() {
	db_auth := fmt.Sprintf("%v:%v@(%v)/%v?parseTime=true&charset=utf8mb4", config.Config.DbUser, config.Config.DbPass, config.Config.DbPort, config.Config.DbName)
	Db, err = sql.Open(config.Config.SQLDriver, db_auth)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY,
		uuid VARCHAR(255),
		nickname VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(20) NOT NULL,
		birthday DATE,
		created_at DATETIME
		)`, tabelNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdC := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(30),
		commitee VARCHAR(20),
		imagepath VARCHAR(30),
		address VARCHAR(50),
		contact VARCHAR(50),
		birthday DATETIME,
		url VARCHAR(80),
		active BOOLEAN NOT NULL,
		created_at DATETIME)`, tableNameCouncilor)

	_, err = Db.Exec(cmdC)

	if err != nil {
		log.Fatalln(err)
	}

	cmdQ := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT AUTO_INCREMENT PRIMARY KEY,
			overview TEXT,
			category VARCHAR(30),
			content TEXT,
			answer TEXT,
			held_time VARCHAR(20),
			councilor_id INTEGER,
			created_at DATETIME)`, tableNameQuestion)

	_, err := Db.Exec(cmdQ)

	if err != nil {
		log.Fatalln(err)
	}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT AUTO_INCREMENT PRIMARY KEY,
			uuid VARCHAR(255) NOT NULL UNIQUE,
			nickname VARCHAR(30),
			user_id INTEGER,
			created_at DATETIME)`, tableNameSession)

	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}

	cmdF := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INTEGER,
			councilor_id INTEGER,
			category TEXT,
			created_at DATETIME)`, tableNameFavorite)

	_, err = Db.Exec(cmdF)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
