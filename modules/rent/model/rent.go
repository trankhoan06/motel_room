package model

import (
	"main.go/modules/upload"
	"main.go/modules/user/model"
	"time"
)

type TypeRoom int

const (
	TypeApartment TypeRoom = iota
	TypeMotelRoom
	TypePrivateHouse
)

type StatusRent int

const (
	StatusDeletedRent StatusRent = iota
	StatusDoingRent
)

type Rent struct {
	Id            int               `json:"id" gorm:"id"`
	UserId        int               `json:"user_id" gorm:"user_id"`
	Owner         *model.SimpleUser `json:"owner" gorm:"foreignkey: UserId;references:Id"`
	Title         string            `json:"title" gorm:"title"`
	Description   string            `json:"description" gorm:"description"`
	Image         *upload.Image     `json:"image" gorm:"image"`
	RoomType      *TypeRoom         `json:"room_type" gorm:"room_type"`
	Price         int               `json:"price" gorm:"price"`
	DepositAmount int               `json:"deposit_amount" gorm:"deposit_amount"`
	Area          int               `json:"area" gorm:"area"`
	Locate        string            `json:"locate" gorm:"locate"`
	Province      string            `json:"province" gorm:"province"`
	AmountRate    int               `json:"amount_rate" gorm:"amount_rate"`
	Rate          float64           `json:"rate" gorm:"rate"`
	Status        *StatusRent       `json:"status" gorm:"status"`
	CreateAt      time.Time         `json:"create_at" gorm:"create_at"`
	UpdateAt      time.Time         `json:"update_at" gorm:"update_at"`
}
type CreateRent struct {
	UserId        int           `json:"-" gorm:"user_id"`
	Title         string        `json:"title" gorm:"title"`
	Description   string        `json:"description" gorm:"description"`
	Image         *upload.Image `json:"image" gorm:"image"`
	RoomType      *TypeRoom     `json:"room_type" gorm:"room_type"`
	Price         int           `json:"price" gorm:"price"`
	DepositAmount int           `json:"deposit_amount" gorm:"deposit_amount"`
	Area          int           `json:"area" gorm:"area"`
	Locate        string        `json:"locate" gorm:"locate"`
	Province      string        `json:"province" gorm:"province"`
}
type UpdateRent struct {
	Id            int           `json:"id" gorm:"id"`
	UserId        int           `json:"-" gorm:"user_id"`
	Title         *string       `json:"title" gorm:"title"`
	Description   *string       `json:"description" gorm:"description"`
	Image         *upload.Image `json:"image" gorm:"image"`
	RoomType      *TypeRoom     `json:"room_type" gorm:"room_type"`
	Price         *int          `json:"price" gorm:"price"`
	DepositAmount *int          `json:"deposit_amount" gorm:"deposit_amount"`
	Area          *int          `json:"area" gorm:"area"`
	Locate        *string       `json:"locate" gorm:"locate"`
}
type SimpleRent struct {
	Id            int               `json:"id" gorm:"id"`
	UserId        int               `json:"user_id" gorm:"user_id"`
	Owner         *model.SimpleUser `json:"owner" gorm:"foreignkey: UserId;references:Id"`
	Title         string            `json:"title" gorm:"title"`
	Description   string            `json:"description" gorm:"description"`
	Image         *upload.Image     `json:"image" gorm:"image"`
	RoomType      *TypeRoom         `json:"room_type" gorm:"room_type"`
	Price         int               `json:"price" gorm:"price"`
	DepositAmount int               `json:"deposit_amount" gorm:"deposit_amount"`
	Area          int               `json:"area" gorm:"area"`
	Locate        string            `json:"locate" gorm:"locate"`
	AmountRate    int               `json:"amount_rate" gorm:"amount_rate"`
	Rate          float64           `json:"rate" gorm:"rate"`
}

func (Rent) TableName() string       { return "rent" }
func (CreateRent) TableName() string { return "rent" }
func (UpdateRent) TableName() string { return "rent" }
func (SimpleRent) TableName() string { return "rent" }
