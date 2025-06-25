package model

import (
	"main.go/modules/notify/model"
	"time"
)

type MappingNotify struct {
	Id             int           `json:"id" gorm:"column:id"`
	UserId         int           `json:"user_id" gorm:"column:user_id"`
	NotificationId int           `json:"notification_id" gorm:"column:notification_id"`
	Notify         *model.Notify `json:"notify" gorm:"foreignkey: NotificationId;references:Id"`
	IsRead         bool          `json:"is_read" gorm:"column:is_read"`
	CreateAt       time.Time     `json:"create_at" gorm:"column:create_at"`
	ReadAt         time.Time     `json:"read_at" gorm:"column:read_at"`
}

type CreateMappingNotify struct {
	UserId         int `json:"user_id" gorm:"column:user_id"`
	NotificationId int `json:"notification_id" gorm:"column:notification_id"`
}

func (MappingNotify) TableName() string       { return "mapping_notify" }
func (CreateMappingNotify) TableName() string { return "mapping_notify" }
