package model

import "time"

type Search struct {
	Id         int       `json:"id" gorm:"id"`
	UserId     int       `json:"user_id" gorm:"user_id"`
	Content    string    `json:"content" gorm:"content"`
	SearchTime int       `json:"search_time" gorm:"search_time"`
	CreateAt   time.Time `json:"create_at" gorm:"create_at"`
}
type SearchRent struct {
	Content string      `json:"content" gorm:"content"`
	Limit   int         `json:"limit"`
	OffSet  int         `json:"offSet"`
	Sort    *SortSearch `json:"sort" gorm:"sort"`
}
type CreateSearch struct {
	UserId  int    `json:"-" gorm:"user_id"`
	Content string `json:"content" gorm:"content"`
}
type UpdateSearch struct {
	UserId int `json:"-" gorm:"user_id"`
	Id     int `json:"id" gorm:"id"`
}
type ListSearch struct {
	UserId  int    `json:"-" gorm:"user_id"`
	Content string `json:"content" gorm:"content"`
	Limit   int    `json:"limit" gorm:"limit"`
}

func (Search) TableName() string       { return "search" }
func (CreateSearch) TableName() string { return "search" }
