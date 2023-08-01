package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Age   int
	Email string
	Work  string
}