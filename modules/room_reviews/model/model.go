package model

import "main.go/modules/user/model"

type Reviews struct {
	Id        int               `json:"id" gorm:"column:id"`
	RentId    int               `json:"rent_id" gorm:"column:rent_id"`
	UserId    int               `json:"user_id" gorm:"column:user_id"`
	User      *model.SimpleUser `json:"user" gorm:"foreignkey: UserId;references:Id"`
	ParentId  int               `json:"parent_id" gorm:"column:parent_id"`
	IsOwner   bool              `json:"is_owner" gorm:"column:is_owner"`
	Comment   string            `json:"comment" gorm:"column:comment"`
	Rate      float64           `json:"rate" gorm:"column:rate"`
	CreatedAt string            `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string            `json:"updated_at" gorm:"column:updated_at"` //ON UPDATE CURRENT_TIMESTAMP
}
type CreateReviews struct {
	Id       int     `json:"-" gorm:"column:id"`
	RentId   int     `json:"rent_id" gorm:"column:rent_id"`
	UserId   int     `json:"-" gorm:"column:user_id"`
	ParentId int     `json:"parent_id" gorm:"column:parent_id"`
	Comment  string  `json:"comment" gorm:"column:comment"`
	Rate     float64 `json:"rate" gorm:"column:rate"`
	IsOwner  bool    `json:"-" gorm:"column:is_owner"`
}
type OwnerResponseReviews struct {
	Id      int    `json:"-" gorm:"column:id"`
	RentId  int    `json:"rent_id" gorm:"column:rent_id"`
	UserId  int    `json:"-" gorm:"column:user_id"`
	Comment string `json:"comment" gorm:"column:comment"`
}
type UpdateReviews struct {
	Id      int      `json:"id" gorm:"column:id"`
	RentId  int      `json:"rent_id" gorm:"column:rent_id"`
	UserId  int      `json:"-" gorm:"column:user_id"`
	Comment *string  `json:"comment" gorm:"column:comment"`
	Rate    *float64 `json:"rate" gorm:"column:rate"`
}
type DeletedReviews struct {
	Id     int `json:"id" gorm:"column:id"`
	RentId int `json:"rent_id" gorm:"column:rent_id"`
	UserId int `json:"-" gorm:"column:user_id"`
}

func (CreateReviews) TableName() string { return "room_reviews" }
func (Reviews) TableName() string       { return "room_reviews" }
func (UpdateReviews) TableName() string { return "room_reviews" }
