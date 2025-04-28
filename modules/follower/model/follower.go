package model

import (
	modelUser "main.go/modules/user/model"
	"time"
)

type Follower struct {
	Id         int                   `json:"id" gorm:"id"`
	FollowerId int                   `json:"follower_id" gorm:"follower_id"`
	UserId     int                   `json:"user_id" gorm:"user_id"`
	Owner      *modelUser.SimpleUser `json:"owner" gorm:"foreignkey:UserId;references:id"`
	CreateAt   time.Time             `json:"create_at" gorm:"create_at"`
}
type CreateFollower struct {
	FollowerId int `json:"follower_id" gorm:"follower_id"`
	UserId     int `json:"user_id" gorm:"user_id"`
}
type SimpleFollower struct {
	Amount int `json:"amount" gorm:"amount"`
}
type SimpleFollowing struct {
	Amount int `json:"amount" gorm:"amount"`
}

func (Follower) TableName() string       { return "follow" }
func (CreateFollower) TableName() string { return "follow" }
