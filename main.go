package main

import (
	"CA_MISSION/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var user model.User
var gacha model.Gacha
var character model.Character
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
  var user model.User
  db := sqlConnect()
  if err := ctx.ShouldBindJSON(&user);
  err != nil {
      ctx.JSON(400, gin.H{"error": err.Error()})
      return
  }
  db.Create(&user)
  ctx.JSON(200, gin.H{
    "token":user.Mail,
  })
  defer db.Close()
}

// Userを表示する関数
func UserGet (ctx *gin.Context){
  var user= model.User{}
  db := sqlConnect()
  token :=ctx.Request.Header.Get("x-token")
  response :=  db.Select("name").Where("mail = ?", token).First(&user)
  ctx.JSON(200, gin.H{
      "data":response,
  })
     defer db.Close()
}

//Userを更新する関数
func UserPut (ctx *gin.Context){
  var user model.User
  db := sqlConnect()
  token :=ctx.Request.Header.Get("x-token")
  if err := ctx.ShouldBindJSON(&user);
  err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
    db.Model(&user).Where("mail=?",token).Update("name", user.Name)
    ctx.JSON(200, gin.H{
        "name":user.Name+"に変更されました",
    })
    defer db.Close()
} 

//ガチャPost関数
func CharaPost (ctx *gin.Context){
  rand.Seed(time.Now().UnixNano())
  // token :=ctx.Request.Header.Get("x-token")
  if err := ctx.ShouldBindJSON(&gacha);
  err != nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  for i:=0;i<gacha.Time;i++{
    if gacha.Time >5{
      // chance :=rand.Intn(100)
      
  //mysql からjsonみたいにむき出す
      
      
    }
  }
}
//characterテーブルからキャラを抽出しuserテーブルに挿入する関数
func GetChara(token string , chance string){
  db := sqlConnect()
  fmt.Println(db.Where("Mail=?",token).Find(&user))
  db.First(&character, token)
  defer db.Close()
}
//キャラクターテーブル生成
func CharaCreate (){
  db := sqlConnect()
  db.AutoMigrate(&model.Character{})
  charaNames :=[]string{"Doragon","Dracula","Witch","Vampire","Ghost"}
  charaChance :=[]string{"20","50","70","90","100"}
  for i :=0;i<len(charaNames);i++{
  charaData := map[string]string{"Name":charaNames[i],"chance":charaChance[i]}
    db.Create(&model.Character{
      Name:charaData["Name"],
      Percent:charaData["chance"],
    })
  }
  defer db.Close()
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
