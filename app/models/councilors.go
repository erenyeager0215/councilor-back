package models

import (
	"log"
	"time"
)



type Councilor struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Commitee  string `json:"commitee"`
	Image string `json:"image"`
	Birthday  time.Time `json:"birthday"`
	Adress    string `json:"address"`
	Contact    string `json:"contact"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"-"`
	// Questions []Question
}



func GetCouncilor(id int)(Councilor ,error){
	var c Councilor
	cmd:= "SELECT id,name,commitee,image,birthday,address,contact,url FROM test_table WHERE id = ?"
	err = Db.QueryRow(cmd,id).Scan(
		&c.Id,
		&c.Name,
		&c.Commitee,
		&c.Image,
		&c.Birthday,
		&c.Adress,
		&c.Contact,	 
		&c.Url,
	)	
	if err != nil{
		log.Fatal(err)
	}
	return c,err
}



func GetCouncilorList()(councilors []Councilor,err error){
	cmd:= "SELECT id,name,commitee,image,birthday,address,contact,url FROM test_table"
	rows,err := Db.Query(cmd)
	if err != nil{
		log.Fatalln(err)		
	}
	for rows.Next(){
		var c Councilor 
		err= rows.Scan(
			&c.Id,
			&c.Name,
			&c.Commitee,
			&c.Image,
			&c.Birthday,
			&c.Adress,
			&c.Contact,
			&c.Url,
		)
		if err != nil{
			log.Fatalln(err)
		}
		councilors = append(councilors,c)
	}
	rows.Close()

	return councilors,err
}