package model

import "time"

type Notify struct {
	Id       int       `json:"id" gorm:"column:id"`
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	Content  string    `json:"content" gorm:"column:content"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type CreateNotify struct {
	Id      int    `json:"-" gorm:"column:id"`
	UserId  int    `json:"-" gorm:"column:user_id"`
	Content string `json:"content" gorm:"column:content"`
}
type UpdateNotify struct {
	Id      int    `json:"id" gorm:"column:id"`
	UserId  int    `json:"-" gorm:"column:user_id"`
	Content string `json:"content" gorm:"column:content"`
}
type DeletedNotify struct {
	Id     int `json:"ud" gorm:"column:id"`
	UserId int `json:"-" gorm:"column:user_id"`
}

func (Notify) TableName() string       { return "notify" }
func (CreateNotify) TableName() string { return "notify" }
func (UpdateNotify) TableName() string { return "notify" }
