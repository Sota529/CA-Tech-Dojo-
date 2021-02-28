package model

import (	
	_"github.com/jinzhu/gorm"
)
type Gacha struct{
	Time int `json:"times"`
	Token string `json:"token"`
}
type Post struct{
	PostID uint `gorm:"not null"`
	CharaID uint
	Chara string 
}
type Result struct {
  CharaID  uint
  Chara string
}

type Character struct {
	ID uint `gorm:"unique;not null;PRIMARY_KEY;autoIncrement"`
  Name string 	`gorm:"not null"`
	Percent string `gorm:"not null"`
}

