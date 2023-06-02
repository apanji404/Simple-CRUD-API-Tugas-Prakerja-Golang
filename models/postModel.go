package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Menu  string
	Price int
}
