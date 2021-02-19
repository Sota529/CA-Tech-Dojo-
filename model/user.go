package model

import (
	// "github.com/jinzhu/gorm"
)
type User struct {
	id uint 
  Name string 	`json:"name"`
	Mail string    `json:"mail"`
}
