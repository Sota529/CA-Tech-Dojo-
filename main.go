package main

import (
	"CA_MISSION/model"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var user model.User
var character model.Character
var gacha model.Gacha
var post = model.Post{}
func main() {
  db := sqlConnect()
  // db.DropTable(&model.User{})//初期化
  db.DropTable(&model.Character{})//初期化
  db.DropTable(&model.Post{})//初期化
  db.AutoMigrate(&model.User{})
  CharaCreate()
  defer db.Close()
  router := gin.Default()
// GETメソッド
router.GET("/user/get",UserGet)
router.GET("/character/list",CharaGet)

// POSTメソッド
router.POST("/user/create", UserPost)
router.POST("/gacha/draw", CharaPost)

// PUTメソッド
router.PUT("/user/update",UserPut)

router.Run()
}

//Userを作成する関数
func UserPost (ctx *gin.Context){
  var user= model.User{}
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
  db.Select("name").Where("mail = ?", token).First(&user)
  ctx.JSON(200, gin.H{
      "Name":user.Name,
  })
     defer db.Close()
}

//Userを更新する関数
func UserPut (ctx *gin.Context){
  var user =model.User{}
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
  // var post = model.Post{}
  token :=ctx.Request.Header.Get("x-token")
  if err := ctx.ShouldBindJSON(&gacha);
  err != nil {
    ctx.JSON(400, gin.H{"error": err.Error()})
    return
  }
  for i:=0;i<gacha.Time;i++{
    rand.Seed(time.Now().UnixNano())
    chance :=(rand.Intn(100))
    GetChara(token ,chance)
  ctx.JSON(200, gin.H{
    "name":post.Chara,
})  
  }
}
//characterテーブルからキャラを抽出しuserテーブルに挿入する関数
func GetChara(token string , chance int){
  var user =model.User{}
  var character =model.Character{}
  db := sqlConnect()
  defer db.Close()
  db.AutoMigrate(&model.Post{})
    for  i:=1;i<=5;i++{
      db.Select("name").Where("id=?", i).Find(&character) 
      db.Select("percent").Where("id=?", i).Find(&character)
      percent,_:=strconv.Atoi(character.Percent)
      if (chance >percent){
        db.Select("id").Where("mail=?",token).First(&user) 
        db.Select("id").Where("percent=?", percent).Find(&character)
        post.CharaID= character.ID
        post.PostID=user.ID
        post.Chara=character.Name
        db.Create(&post)
        db.Model(&user).Select("*").Joins("left join posts on posts.post_id = users.id")     
      }
      if (chance >percent){
      break
      }}
}
//charaGet関数
func CharaGet(ctx *gin.Context){
  var user= model.User{}
  var result  []model.Result
  db := sqlConnect()
  defer db.Close()
  token :=ctx.Request.Header.Get("x-token")
  db.Select("posts.chara_id,posts.chara").Where("mail = ?", token).Joins("left join posts on posts.post_id = users.id").Find(&user).Scan(&result)
  
  ctx.JSON(200, gin.H{
      "characters":result,
  })
}
//character生成
func CharaCreate (){
  db := sqlConnect()
  db.AutoMigrate(&model.Character{})
  charaNames :=[]string{"Doragon","Dracula","Witch","Vampire","Ghost"}
  charaChance :=[]string{"100","90","70","40","0"}
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
