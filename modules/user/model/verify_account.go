package model

import "time"

type VerifyAccount struct {
	Id               int       `json:"id" gorm:"id"`
	UserId           int       `json:"user_id" gorm:"user_id"`
	Code             int       `json:"code" gorm:"code"`
	Verify           bool      `json:"verify" gorm:"verify"`
	IsVerifyEmail    bool      `json:"is_verify_email" gorm:"is_verify_email"`
	IsForgotPassword bool      `json:"is_forgot_password" gorm:"is_forgot_password"`
	CreateAt         time.Time `json:"create_at" gorm:"create_at"`
	Expire           time.Time `json:"expire" gorm:"expire"`
}
type CreateVerifyAccount struct {
	UserId           int       `json:"user_id" gorm:"user_id"`
	IsVerifyEmail    bool      `json:"is_verify_email" gorm:"is_verify_email"`
	IsForgotPassword bool      `json:"is_forgot_password" gorm:"is_forgot_password"`
	Code             int       `json:"code" gorm:"code"`
	Expire           time.Time `json:"expire" gorm:"expire"`
}
type VerifyAccountCode struct {
	Email            string `json:"email" gorm:"email"`
	Code             int    `json:"code" gorm:"code"`
	IsVerifyEmail    bool   `json:"is_verify_email" gorm:"is_verify_email"`
	IsForgotPassword bool   `json:"is_forgot_password" gorm:"is_forgot_password"`
}

func (CreateVerifyAccount) TableName() string { return "send_code" }
func (VerifyAccount) TableName() string       { return "send_code" }
