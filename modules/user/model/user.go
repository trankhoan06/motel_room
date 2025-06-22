package model

import (
	"main.go/modules/upload"
	"time"
)

type RoleUser int

const (
	RoleUserUser RoleUser = iota
	RoleUserHost
	RoleUserAdmin
)

type StatusUser int

const (
	StatusUserDoing StatusUser = iota
	StatusUserDeleted
)

type User struct {
	Id              int           `json:"id" gorm:"column:id"`
	Email           string        `json:"email" gorm:"column:email"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Image           *upload.Image `json:"image" gorm:"column:image"`
	Password        string        `json:"-" gorm:"column:password"`
	FistName        string        `json:"fist_name" gorm:"column:fist_name"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            *RoleUser     `json:"role" gorm:"column:role"`
	AmountFollower  int           `json:"amount_follower" gorm:"column:amount_follower"`
	AmountFollowing int           `json:"amount_following" gorm:"column:amount_following"`
	Address         string        `json:"address" gorm:"column:address"`
	Status          *StatusUser   `json:"status" gorm:"column:status"`
	IsEMail         bool          `json:"is_email" gorm:"column:is_email"`
	CreatedAt       time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time     `json:"updated_at" gorm:"column:updated_at"`
	//	ON UPDATE CURRENT_TIMESTAMP
}
type Login struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
type Register struct {
	Id       int           `json:"-" gorm:"column:id"`
	Email    string        `json:"email" gorm:"column:email"`
	Salt     string        `json:"-" gorm:"column:salt"`
	Password string        `json:"password" gorm:"column:password"`
	FistName string        `json:"fist_name" gorm:"column:fist_name"`
	LastName string        `json:"last_name" gorm:"column:last_name"`
	Image    *upload.Image `json:"image" gorm:"column:image"`
	Phone    string        `json:"phone" gorm:"column:phone"`
	Address  string        `json:"address" gorm:"column:address"`
	IsEMail  bool          `json:"-" gorm:"column:is_email"`
}

type UpdateUser struct {
	Email    string        `json:"-" gorm:"column:email"`
	FistName *string       `json:"fist_name" gorm:"column:fist_name"`
	Image    *upload.Image `json:"image" gorm:"column:image"`
	LastName *string       `json:"last_name" gorm:"column:last_name"`
	Phone    *string       `json:"phone" gorm:"column:phone"`
	Address  *string       `json:"address" gorm:"column:address"`
}
type ChangePassword struct {
	Id          int    `json:"-" gorm:"column:id"`
	Email       string `json:"-" gorm:"column:email"`
	OldPassword string `json:"old_password" gorm:"column:old_password"`
	NewPassword string `json:"new_password" gorm:"column:new_password"`
}
type ForgotPassword struct {
	Email string `json:"email" gorm:"column:email"`
}
type NewPasswordForgot struct {
	Email       string `json:"email" gorm:"column:email"`
	Token       string `json:"token" gorm:"token"`
	NewPassword string `json:"new_password" gorm:"column:new_password"`
}
type LoginMedia struct {
	YourAccessToken string `json:"your_access_token" gorm:"column:your_access_token"`
}
type StatusUserValue struct {
	Status *StatusUser `json:"status" gorm:"column:status"`
}
type SimpleUser struct {
	Id       int           `json:"-" gorm:"column:id"`
	Image    *upload.Image `json:"image" gorm:"column:image"`
	FistName string        `json:"fist_name" gorm:"column:fist_name"`
	LastName string        `json:"last_name" gorm:"column:last_name"`
	Address  string        `json:"address" gorm:"column:address"`
}

func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetRole() *RoleUser {
	return u.Role
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u User) TableName() string       { return "user" }
func (u SimpleUser) TableName() string { return "user" }
func (u UpdateUser) TableName() string { return "user" }
func (u Login) TableName() string      { return "user" }
func (u Register) TableName() string   { return "user" }
