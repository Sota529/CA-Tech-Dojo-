package model

import (
	_"github.com/jinzhu/gorm"
)
type Character struct {
	CharacterID uint `gorm:"unique;not null;PRIMARY_KEY"`
  Name string 	`json:"name" ;gorm:";not null" `
	Gender string    `json:"json";gorm:";not null" `
}