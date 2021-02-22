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
  db.DropTable(&model.User{})//初期化
  db.DropTable(&model.Character{})//初期化
  db.AutoMigrate(&model.User{})
  CharaCreate()
  defer db.Close()
  router := gin.Default()

// GETメソッド
router.GET("/user/get",UserGet)

// POSTメソッド
router.POST("/user/create", UserPost)
router.POST("/gacha/draw", CharaPost)

// PUTメソッド
router.PUT("/user/update",UserPut)
  
router.Run()
}

//Userを作成する関数
func UserPost (ctx *gin.Context){
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
}

// Userを表示する関数
func UserGet (ctx *gin.Context){
  var json model.User
  db := sqlConnect()
  token :=ctx.Request.Header.Get("x-token")
  response :=db.Where("Mail=?",token).Find(&json)
  ctx.JSON(200, gin.H{
      "data":response,
  })
     defer db.Close()
}

//Userを更新する関数
func UserPut (ctx *gin.Context){
  var json model.User
  db := sqlConnect()
  token :=ctx.Request.Header.Get("x-token")
  if err := ctx.ShouldBindJSON(&json);
  err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Model(&json).Where("mail=?",token).Update("name", json.Name)
    ctx.JSON(200, gin.H{
        "name":json.Name+"に変更されました",
    })
    defer db.Close()
} 

//ガチャPost関数
func CharaPost (ctx *gin.Context){
  //userテーブル
  var json model.User
  //ガチャテーブル
  var gacha model.Gacha
  rand.Seed(time.Now().UnixNano())
  if err := ctx.ShouldBindJSON(&gacha);
  err != nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  db := sqlConnect()
  for i:=0;i<gacha.Time;i++{
    if gacha.Time >5{
      fmt.Println(db.Model(json).Select("*").Joins("left join gacha on user.mail = gacha.token").Scan(&gacha))
      
      
    }
    fmt.Println((db.Exec("SELECT name FROM users WHERE name = ?",1)))
    rand.Intn(100)
  }
  
  
  
  defer db.Close()
}

//キャラクターテーブル生成
func CharaCreate (){
  db := sqlConnect()
  db.AutoMigrate(&model.Character{})
  charaName := []string{ "Doragon", "Dracula","Witch","Vampire","Ghost"}
  for i :=0;i<len(charaName);i++{
    db.Create(&model.Character{
      Name:charaName[i],
    })
  }
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
