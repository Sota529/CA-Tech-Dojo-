package model

import (	
	_"github.com/jinzhu/gorm"
)
type Gacha struct{
	Time int `json:"times"`
	Token string `json:"token"`
}

type Character struct {
	CharacterID uint `gorm:"unique;not null;PRIMARY_KEY;autoIncrement"`
  Name string 	`gorm:"not null;PRIMARY_KEY" `
}

