package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
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

type Book struct {
	Title     string `json:"title" form:"title" query:"title"`
	FirstName string `json:"firstname"  form:"firstname" query:"firstname"`
	LastName  string `json:"lastname"  form:"lastname" query:"lastname"`
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

	e.GET("/show", show)
	e.GET("/councilors",getCouncilors)
	e.GET("/councilor/:id",getCouncilor)
	e.POST("/save", save)
	e.POST("/save1", save1)
	e.POST("/books", books)
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

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

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

// http://localhost:1323/show?team=x-men&member=wolverine
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}



// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	
	u := &User{}
	u.Name = name
	u.Email = email

	Db, _ := sql.Open("sqlite3", "./test.sql")
	defer Db.Close()

	cmd := "INSERT INTO test_table(name,email)VALUES(?,?)"
	_, err := Db.Exec(cmd, u.Name, u.Email)
	if err != nil {
		log.Fatalln(err)
	}

	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func save1(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

// e.POST("/books", books)
// formで受けたデータをJSONに変換する
func books(c echo.Context) error {
	name := c.FormValue("title")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	b := &Book{}
	b.Title = name
	b.FirstName = firstName
	b.LastName = lastName

	if err := c.Bind(b); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, b)
}

// return c.String(http.StatusOK, "name:"+name+", first:"+firstName+",lastName:"+lastName)
