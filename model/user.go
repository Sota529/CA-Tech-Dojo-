package model

import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
  Name string 	`json:"name"`
	Mail string    `json:"mail"`
}
