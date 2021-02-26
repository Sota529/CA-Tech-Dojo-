package model

import (	
	_"github.com/jinzhu/gorm"
)
type Gacha struct{
	Time int `json:"times"`
	Token string `json:"token"`
}
type Post struct{
	ID uint `gorm:"not null"`
	Chara string 
}

type Character struct {
	ID uint `gorm:"unique;not null;PRIMARY_KEY;autoIncrement"`
  Name string 	`gorm:"not null;PRIMARY_KEY"`
	Percent string `gorm:"not null"`
}

