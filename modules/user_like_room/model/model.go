package model

import (
	modelRent "main.go/modules/rent/model"
	"time"
)

type StatusUserLikeRoom int

const (
	StatusDeletedUserLikeRoom StatusUserLikeRoom = iota
	StatusDoingUserLikeRoom
)

type UserLikeRoom struct {
	Id         int                   `json:"id" gorm:"column:id"`
	UserId     int                   `json:"user_id" gorm:"column:user_id"`
	RentId     int                   `json:"rent_id" gorm:"column:rent_id"`
	SimpleRent *modelRent.SimpleRent `json:"simple_rent" gorm:"foreignkey:RentId;references:Id""`
	Status     *StatusUserLikeRoom   `json:"status" gorm:"column:status"`
	CreatedAt  time.Time             `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time             `json:"updated_at" gorm:"column:updated_at"`
}
type CreateUserLikeRoom struct {
	UserId int `json:"user_id" gorm:"column:user_id"`
	RentId int `json:"rent_id" gorm:"column:rent_id"`
}

func (UserLikeRoom) TableName() string { return "user_like_room" }
