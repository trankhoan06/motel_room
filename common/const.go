package common

import "errors"

const Current_user = "Current_user"

var (
	ErrVerifyEmail      = errors.New("account don't verify")
	ErrVerifyCode       = errors.New("code wrong")
	ErrVerifyCodeExpire = errors.New("code expire")
	ErrVerify           = errors.New("you don't verify code")
	ErrForgotPassword   = errors.New("expire change forgot password")
	ErrLogin            = errors.New("you need login")
	ErrEmailRequire     = errors.New("email is require")
)

type TreeReview struct {
	Val   interface{}
	Child *[]TreeReview
}

func NewTreeReview(val interface{}) *TreeReview {
	return &TreeReview{
		Val:   val,
		Child: &[]TreeReview{},
	}
}
