package main

import (
	"fmt"
	"CA_MISSION/model"

	"github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)



func main() {
  db := sqlConnect()
  fmt.Println(db.HasTable("users"))
  db.AutoMigrate(&model.User{},)
//   db.CreateTable(&model.User{})
  
  //   if db.HasTable("users") == false {
    // db.CreateTable(&model.User{})
//   }
  defer db.Close()
  router := gin.Default()

// GETメソッド
  router.GET("/user/get", func(ctx *gin.Context){
    var json model.User
    db := sqlConnect()
    ctx.JSON(200, gin.H{
        "data":db.First(&json),
    })
    fmt.Println("user got!")
    defer db.Close()
})
// POSTメソッド
  router.POST("/user/create", func(ctx *gin.Context){
    var json model.User
    db := sqlConnect()
    db.Select("Name").Create(&json)
    if err := ctx.ShouldBindJSON(&json); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(200, gin.H{
      "token":json.Name,
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
 fmt.Println("接続開始")
  db, err := gorm.Open("mysql", "go_test:password@tcp(db:3306)/go_database?charset=utf8&parseTime=True&loc=Local")
if err != nil {
    panic(err)
}
fmt.Println("接続成功")
  return db
}
