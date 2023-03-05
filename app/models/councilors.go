package models

import (
	"log"
	"time"
)

type Councilor struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Commitee  string    `json:"commitee"`
	ImagePath string    `json:"imagepath"`
	Birthday  time.Time `json:"birthday"`
	Adress    string    `json:"address"`
	Contact   string    `json:"contact"`
	Url       string    `json:"url"`
	Active    bool      `json:"-"`
	CreatedAt time.Time `json:"-"`
	// Questions []Question
}

type CouncilorsRanking struct {
	CouncilorName string `json:"name"`
	Score         int    `json:"score"`
}

func GetCouncilor(id int) (Councilor, error) {
	var c Councilor
	cmd := "SELECT id,name,commitee,imagepath,birthday,address,contact,url FROM councilors WHERE id = ?"
	err = Db.QueryRow(cmd, id).Scan(
		&c.Id,
		&c.Name,
		&c.Commitee,
		&c.ImagePath,
		&c.Birthday,
		&c.Adress,
		&c.Contact,
		&c.Url,
	)
	if err != nil {
		log.Fatal(err)
	}
	return c, err
}

func GetCouncilorList() (councilors []Councilor, err error) {
	cmd := "SELECT id,name,commitee,imagepath,address,contact,birthday,url FROM councilors"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var c Councilor
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Commitee,
			&c.ImagePath,
			&c.Adress,
			&c.Contact,
			&c.Birthday,
			&c.Url,
		)
		if err != nil {
			log.Fatalln(err)
		}
		councilors = append(councilors, c)
	}
	rows.Close()

	return councilors, err
}

func GetTopFiveOfCouncilors() (councilorsRanking []CouncilorsRanking, err error) {
	cmd := "SELECT councilors.name , COUNT(*) AS score FROM favorite JOIN councilors ON favorite.councilor_id = councilors.id GROUP BY councilors.name ORDER BY 2 DESC LIMIT 5"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var cr CouncilorsRanking
		err = rows.Scan(
			&cr.CouncilorName,
			&cr.Score,
		)
		if err != nil {
			log.Fatalln(err)
		}
		councilorsRanking = append(councilorsRanking, cr)
	}
	return councilorsRanking, err
}
