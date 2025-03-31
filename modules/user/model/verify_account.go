package model

import "time"

type VerifyAccount struct {
	Id       int       `json:"id" gorm:"id"`
	UserId   int       `json:"user_id" gorm:"user_id"`
	Token    string    `json:"token" gorm:"token"`
	Code     int       `json:"code" gorm:"code"`
	Verify   bool      `json:"verify" gorm:"verify"`
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
	Expire   time.Time `json:"expire" gorm:"expire"`
}
type CreateVerifyAccount struct {
	UserId int       `json:"user_id" gorm:"user_id"`
	Token  string    `json:"token" gorm:"token"`
	Code   int       `json:"code" gorm:"code"`
	Expire time.Time `json:"expire" gorm:"expire"`
}
type VerifyAccountCode struct {
	Email string `json:"email" gorm:"email"`
	Token string `json:"token" gorm:"token"`
	Code  int    `json:"code" gorm:"code"`
}
type VerifyToken struct {
	Token   string `json:"token" gorm:"token"`
	Email   string `json:"email" gorm:"email"`
	IsLogin bool   `json:"is_login" gorm:"is_login"`
}

func (CreateVerifyAccount) TableName() string { return "send_code" }
func (VerifyAccount) TableName() string       { return "send_code" }
