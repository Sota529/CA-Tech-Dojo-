package main

import (
	"fmt"
	"CA_MISSION/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)



func main() {
  fmt.Println("Ha")
  db := sqlConnect()
  defer db.Close()

  router := gin.Default()

// GETメソッド
  router.GET("/user/get", func(ctx *gin.Context){
    fmt.Println("user got!")
})
// POSTメソッド
  router.POST("/user/create", func(ctx *gin.Context){
    db := sqlConnect()
    id :=ctx.PostForm("id")
    name := ctx.PostForm("name")
    fmt.Println(name)
    mail := ctx.PostForm("mail")
    fmt.Println(name +"user Created!"+mail+"Mail")
    db.Create(&model.User{Name: name, Mail: mail})
    ctx.JSON(200, gin.H{
      "token":id,
    })
    defer db.Close()
})
// PUTメソッド
  router.PUT("/user/update", func(ctx *gin.Context){
    fmt.Println("user updated!")
})
  
  router.Run()
}

// mysql接続関数
func sqlConnect() (database *gorm.DB) {
  db, err := gorm.Open("mysql", "go_test:password@tcp(db:3306)/go_database?charset=utf8&parseTime=True&loc=Local")
if err != nil {
    panic(err)
}
defer func() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}()
db.LogMode(true)
if err := db.DB().Ping(); err != nil {
    panic(err)
}

db.AutoMigrate(
    &model.User{},
)
  return db
}



