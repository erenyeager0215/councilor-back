package models

import (
	"log"
	"time"
)

type CouncilorList struct {
	Councilor []Councilor
}

type Councilor struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Commitee  string `json:"commitee"`
	ImagePath string `json:"iamge"`
	Birthday  time.Time `json:"birthday"`
	Adress    string `json:"address"`
	TelNum    string `json:"tel"`
	CreatedAt time.Time `json:"created_at"`
	// Questions []Question
}


func GetCouncilor(id int)(c *Councilor,err error){
	c = new(Councilor)
	cmd:= "SELECT * FROM councils WHERE id = ?"
	err = Db.QueryRow(cmd,id).Scan(
		c.Id,
		c.Name,
		c.Commitee,
		c.ImagePath,
		c.Birthday,
		c.Adress,
		c.TelNum,
		c.CreatedAt, 
	)
	if err != nil{
		log.Fatal(err)
	}
	return c,err
}