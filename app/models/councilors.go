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
	Age_under19   int    `json:"under19"`
	Age_20s       int    `json:"age_20s"`
	Age_30s       int    `json:"age_30s"`
	Age_40s       int    `json:"age_40s"`
	Age_50s       int    `json:"age_50s"`
	Age_60s       int    `json:"age_60s"`
	Age_70s       int    `json:"age_70s"`
	Age_over80    int    `json:"over80"`
	Total         int    `json:"total"`
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
	cmd := "SELECT councilors.name,SUM(CASE WHEN age <= 19 THEN 1 ELSE 0 END) AS '19~',SUM(CASE WHEN age BETWEEN 20 AND 29 THEN 1 ELSE 0 END) AS '20~29',SUM(CASE WHEN age BETWEEN 30 AND 39 THEN 1 ELSE 0 END) AS '30~39',SUM(CASE WHEN age BETWEEN 40 AND 49 THEN 1 ELSE 0 END) AS '40~49',SUM(CASE WHEN age BETWEEN 50 AND 59 THEN 1 ELSE 0 END) AS '50~59',SUM(CASE WHEN age BETWEEN 60 AND 69 THEN 1 ELSE 0 END) AS '60~69',SUM(CASE WHEN age BETWEEN 70 AND 79 THEN 1 ELSE 0 END) AS '70~79',SUM(CASE WHEN age >= 80 THEN 1 ELSE 0 END) AS '80~', COUNT(*) AS 'total' FROM favorite JOIN users ON favorite.user_id = users.id JOIN councilors ON favorite.councilor_id = councilors.id,(SELECT id, TIMESTAMPDIFF(YEAR, `birthday`, CURDATE()) AS age FROM users) AS age_table WHERE users.id = age_table.id GROUP BY councilors.name ORDER BY 10 DESC LIMIT 5"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var cr CouncilorsRanking
		err = rows.Scan(
			&cr.CouncilorName,
			&cr.Age_under19,
			&cr.Age_20s,
			&cr.Age_30s,
			&cr.Age_40s,
			&cr.Age_50s,
			&cr.Age_60s,
			&cr.Age_70s,
			&cr.Age_over80,
			&cr.Total,
		)
		if err != nil {
			log.Fatalln(err)
		}
		councilorsRanking = append(councilorsRanking, cr)
	}
	return councilorsRanking, err
}
