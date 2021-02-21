package main

import (
	"fmt"
  "time"
  "math/rand"
	"CA_MISSION/model"

	"github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
  db := sqlConnect()
  db.DropTable(&model.User{})
  db.AutoMigrate(&model.User{},)
  defer db.Close()
  router := gin.Default()

// GETメソッド
  router.GET("/user/get", func(ctx *gin.Context){
    var json model.User
    db := sqlConnect()
    token :=ctx.Request.Header.Get("x-token")
    response :=db.Where("Mail=?",token).Find(&json)
    ctx.JSON(200, gin.H{
        "data":response,
    })
       defer db.Close()
})
// POSTメソッド
//--------------user--------------------------------//
  router.POST("/user/create", func(ctx *gin.Context){
    var json model.User
    db := sqlConnect()
    if err := ctx.ShouldBindJSON(&json);
    err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Create(&json)
    ctx.JSON(200, gin.H{
      "token":json.Mail,
    })
    defer db.Close()
})
//-----------------user-----------------------//
//-----------------gacha-----------------------//
router.POST("/gacha/draw", func(ctx *gin.Context){
  // var json model.Character
  var gacha model.Gacha
  if err := ctx.ShouldBindJSON(&gacha);
  err != nil {
      ctx.JSON(400, gin.H{"error": err.Error()})
      return
  }
  //rand.Intn()にはつくったキャラの数をいれる。
  rand.Seed(time.Now().UnixNano())
  for i:=0;i<gacha.Time;i++{
    fmt.Println(rand.Intn(5))
  }
  // db := sqlConnect()
  // if err := ctx.ShouldBindJSON(&json);
  // err != nil {
  //     ctx.JSON(400, gin.H{"error": err.Error()})
  //     return
  // }
  // db.Select("CharacterID","Name").Create(&json)
  // ctx.JSON(200, gin.H{
  //   "characterID":json.CharacterID,
  //   "name":json.Name,
  // })
  // defer db.Close()
})
//-----------------gacha-----------------------//
// PUTメソッド
router.PUT("/user/update", func(ctx *gin.Context){
  var json model.User
  db := sqlConnect()
  token :=ctx.Request.Header.Get("x-token")
  if err := ctx.ShouldBindJSON(&json);
  err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Model(&json).Update("name", json.Name).Where("Mail=?",token)
    db.Save(&json)
    ctx.JSON(200, gin.H{
        "name":json.Name+"に変更されました",
    })
    defer db.Close()
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
