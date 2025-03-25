package common

import "errors"

const Current_user = "Current_user"

var (
	ErrVerifyEmail      = errors.New("account don't verify")
	ErrVerifyCode       = errors.New("code wrong")
	ErrVerifyCodeExpire = errors.New("code expire")
)
