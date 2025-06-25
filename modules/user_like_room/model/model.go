package model

import (
	modelUser "main.go/modules/user/model"
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
	SimpleUser *modelUser.SimpleUser `json:"simple_user" gorm:"foreignkey:UserId;references:Id"`
	Status     *StatusUserLikeRoom   `json:"status" gorm:"column:status"`
	CreatedAt  *time.Time            `json:"created_at" gorm:"column:create_at"`
	UpdatedAt  *time.Time            `json:"updated_at" gorm:"column:update_at"`
}
type CreateUserLikeRoom struct {
	UserId int `json:"-" gorm:"column:user_id"`
	RentId int `json:"rent_id" gorm:"column:rent_id"`
}
type ListUserLikeRoom struct {
	RentId int `json:"rent_id" gorm:"column:rent_id"`
}

func (UserLikeRoom) TableName() string       { return "user_like_room" }
func (CreateUserLikeRoom) TableName() string { return "user_like_room" }
