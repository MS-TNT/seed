package models

import "time"

type User struct{
	Id int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name               string `gorm:"column:name;size:255" json:"name"`
	PasswordHash string `gorm:"column:password_hash;size:255" json:"password_hash"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (u *User) TableName() string {
	return "users"
}
