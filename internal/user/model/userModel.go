package userModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Prefix   string
	Phone    string
	Address  string
	Password string
	FkRole   int
}
