package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid      string
	Status   int8   `gorm:"not null;size:8"`
	Username string `gorm:"not null"`
	Nickname string
	Email    string
	Mobile   string
	Identity int8 `gorm:"not null;size:8"`
}
