package main

import "myapp/app/controllers"

func main() {
	// fmt.Println(models.Db)
 controllers.StartMainServer()


// cmd:= `INSERT INTO questions(
// 	overview,
// 	category,
// 	content,
// 	answer,
// 	held_time,
// 	councilor_id
// 	) VALUES(
// 	"テストについて",
// 	"テストカテゴリ",
// 	"テストです",
// 	"市長公室長 テストに引続き行う。",
// 	"令和4年6月定例会",
// 	2
// )`
// 	v,err:= models.Db.Exec(cmd)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	fmt.Println(v)	

// user:= new(models.User)
// user.NickName= "test2"
// user.PassWord="1"
// user.CreateUser()
	
	// fmt.Println(models.Encrypt("1"))
}
