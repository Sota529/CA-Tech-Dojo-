package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Name string
    Email string
  }

func main() {
  db := sqlConnect()
  db.AutoMigrate(&User{})
  defer db.Close()

  router := gin.Default()

// GETメソッド
  router.GET("/user/get", func(ctx *gin.Context){
    fmt.Println("user got!")
})
// POSTメソッド
  router.POST("/user/create", func(ctx *gin.Context){
    fmt.Println("user created!")
})

  router.Run()
}

// mysql接続関数
func sqlConnect() (database *gorm.DB) {
  DBMS := "mysql"
  USER := "go_test"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

  count := 0
  db, err := gorm.Open(DBMS, CONNECT)
  
//   30秒接続できなかったら接続失敗と出る処理
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      fmt.Print(".")
      time.Sleep(time.Second)
      count++
      if count > 30 {
        fmt.Println("")
        fmt.Println("DB接続失敗")
        panic(err)
      }
      db, err = gorm.Open(DBMS, CONNECT)
    }
  }
  fmt.Println("DB接続成功")

  return db
}