package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	utils "myapp/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	utils.LoggingSettings("./test.log")
}

// sqlのDBを指すポインタとしてDbを宣言
var Db *sql.DB

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type Councilor struct{
	Id int `json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
	Address string `json:"address" form:"address" query:"address"`
}

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		// 同時に1つのプログラムの流れのみが資源を占有し、他の使用を排除する
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func main() {
	e := echo.New()

	// Debug mode
	e.Debug = true

	//----------------------------------------------------------------------------
	// Custom middleware
	//----------------------------------------------------------------------------
	// Stats
	s := NewStats()
	//Use関数でリクエストが飛んできたときに事前に決められた共通の処理ができる
	//Proccessはリクエスト情報を取得する
	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// Server header
	e.Use(ServerHeader)

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	//log
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
    //     return func(c echo.Context) error {
    //         if err := next(c); err != nil {
    //             c.Error(err)
    //         }

    //         req_header := fmt.Sprintf("%#v", c.Request().Header)
    //         req_body := fmt.Sprintf("%#v", c.Request().Body)

    //         os.Stdout.Write([]byte(req_header + "\n"))
    //         os.Stdout.Write([]byte(req_body + "\n"))
    //         return next(c)
    //     }
    // })

	//----------------------------------------------------------------------------
	// Custom middleware
	//----------------------------------------------------------------------------

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	
	e.GET("/councilors",getCouncilors)
	e.GET("/councilor/:id",getCouncilor)
	e.POST("/login",login)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		log.Print(u)

		Db, _ := sql.Open("sqlite3", "./test.sql")
		defer Db.Close()

	cmd := "INSERT INTO test_table(name,email)VALUES(?,?)"
	_, err := Db.Exec(cmd, u.Name, u.Email)
	if err != nil {
		log.Fatalln(err)
	}
		return c.JSON(http.StatusCreated, u)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}


	// ----------------------------------------------------------
	// login機能実装
	// ----------------------------------------------------------
	func getUserByNickName(c echo.Context) (User,error){
		u:= new(User)
		var user User
		if err:= c.Bind(u); err != nil{
			log.Fatal(err)
		}
		Db,_:= sql.Open("sqlite3","./test.sql")
		defer Db.Close()

		cmd:= "SELECT * from test_table WHERE name = ?"
		err:= Db.QueryRow(cmd,u.Name).Scan(
			&user.Name,
			&user.Email,
		)
		if err != nil {
			log.Fatal(err)
		}
		return user,err
	}
	
	func login(c echo.Context)error{
		u:= new(User)
		user,err:= getUserByNickName(c)
		log.Println(user)
		if err != nil{
			log.Fatal(err)
		}
		if err= c.Bind(u); err != nil{
			log.Fatal(err)
		}
		if user.Email == u.Email{
			return c.JSON(http.StatusCreated, user)
		}else{
			return c.JSON(http.StatusCreated,"NotFound")
		}
	}

	// ----------------------------------------------------------
	// login機能実装
	// ----------------------------------------------------------


func getCouncilor(c echo.Context)error{
	var councilor Councilor
	id:= c.Param("id")
	log.Print(id)
	Db,_:= sql.Open("sqlite3","././coucils.sql")
	defer Db.Close()
	cmd:= "SELECT * FROM councils WHERE id = ?"
	err := Db.QueryRow(cmd,id).Scan(
		&councilor.Id,
		&councilor.Name,
		&councilor.Address,
	)
	if err != nil{
		log.Fatal(err)
	}
	
	return c.JSON(http.StatusCreated, councilor)
}

func getCouncilors(c echo.Context)error{
	Db, _ := sql.Open("sqlite3", "./coucils.sql")
	defer Db.Close()
	cmd:= "SELECT * FROM councils"
	rows,err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	var councilors []Councilor
	for rows.Next(){
		var councilor Councilor
		err = rows.Scan(
			&councilor.Id,
			&councilor.Name,
			&councilor.Address,
		)
		if err != nil{
			log.Fatal(err)
		}
		councilors = append(councilors,councilor)
	}
	return c.JSON(http.StatusCreated, councilors)
}



