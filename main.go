package main

import (
	"fmt"
	// "time"

	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
    fmt.Println("DB接続開始")

    _, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println(err)
        fmt.Println("DB接続成功")
    }
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "root"
    PROTOCOL := "tcp(docker_MySQL:3306)"
    DBNAME := "go_database"


    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}
   

