package model

import "time"

type Search struct {
	Id         int       `json:"id" gorm:"id"`
	UserId     int       `json:"user_id" gorm:"user_id"`
	Content    string    `json:"content" gorm:"content"`
	SearchTime int       `json:"search_time" gorm:"search_time"`
	CreateAt   time.Time `json:"create_at" gorm:"create_at"`
}
type CreateSearch struct {
	UserId     int    `json:"-" gorm:"user_id"`
	Content    string `json:"content" gorm:"content"`
	SearchTime int    `json:"search_time" gorm:"search_time"`
}

func (Search) TableName() string       { return "search" }
func (CreateSearch) TableName() string { return "search" }
